package translator

import (
	"context"
	"fmt"
	"math"
	nUrl "net/url"
	"strings"
	"url-shortener/internal/domain/model"
	"url-shortener/internal/domain/repo"
)

type UrlTranslator struct {
	linkService repo.LinkService
}

func New(linkService repo.LinkService) *UrlTranslator {
	return &UrlTranslator{
		linkService: linkService,
	}
}

var (
	alphabet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	numeralSystem = uint(len(alphabet))
	shortDomain   = "https://igor.r"
)

func (h *UrlTranslator) ShortenUrl(ctx context.Context, url string) (string, error) {
	_, err := nUrl.ParseRequestURI(url)
	if err != nil {
		return "", err
	}

	err = h.linkService.Insert(ctx, &model.Link{Url: url})
	if err != nil {
		return "", err
	}

	link, err := h.linkService.Select(ctx, &model.Link{Url: url})
	if err != nil {
		return "", err
	}

	shortUrl := h.getShortLinkById(link.ID)
	return shortUrl, nil
}

func (h *UrlTranslator) ExtendUrl(ctx context.Context, url string) (string, error) {
	urlToId := url[len(shortDomain)+1:]

	id := h.getIdByLink(urlToId)

	link, err := h.linkService.Select(ctx, &model.Link{ID: id})
	if err != nil {
		return "", err
	}

	return link.Url, nil
}

func (h *UrlTranslator) getShortLinkById(id uint) string {
	shortUrl := ""

	for id != 0 {
		digit := id % numeralSystem
		shortUrl = string(alphabet[digit]) + shortUrl
		id = uint(math.Floor(float64(id) / float64(numeralSystem)))
	}

	return fmt.Sprintf("%s/%s", shortDomain, shortUrl)
}

func (h *UrlTranslator) getIdByLink(url string) uint {
	id := uint(0)

	for i := 0; i < len(url); i++ {
		char := string(url[len(url)-i-1])
		charValue := strings.Index(alphabet, char)
		id += uint(charValue) * uint(math.Pow(float64(numeralSystem), float64(i)))
	}

	return id
}
