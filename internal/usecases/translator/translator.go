package translator

type UrlTranslator struct {
}

func New() *UrlTranslator {
	return &UrlTranslator{}
}

func (h *UrlTranslator) ShortenUrl(url string) (string, error) {
	return "method Shorten() is not implemented", nil
}

func (h *UrlTranslator) ExtendUrl(url string) (string, error) {
	return "method Extend() is not implemented", nil
}
