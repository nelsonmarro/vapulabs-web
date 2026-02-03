package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nelsonmarro/vapulabs-web/templates/components/products"
	"github.com/nelsonmarro/vapulabs-web/templates/pages"
	"github.com/starfederation/datastar-go/datastar"
)

// Pricing Constants
const (
	VerithBasicMonthlyURL = "https://naphsoft.lemonsqueezy.com/checkout/buy/4e2796d6-5976-4613-a096-a1e86d3dda65?enabled=1240057"
	VerithBasicAnnualURL  = "https://naphsoft.lemonsqueezy.com/checkout/buy/b3497c54-7537-405f-a8b7-93df23abfbc7?enabled=1240020"
	VerithPymeMonthlyURL  = "https://naphsoft.lemonsqueezy.com/checkout/buy/ca111c39-dd50-481b-bb4e-8b3ec82c8690?enabled=1240479"
	VerithPymeAnnualURL   = "https://naphsoft.lemonsqueezy.com/checkout/buy/117a8e15-43a0-41f8-bd5e-6fefa01e9da2?enabled=1240480"
	VerithCorpMonthlyURL  = "https://naphsoft.lemonsqueezy.com/checkout/buy/4d70c256-8929-42b0-ad65-5a9248fd7517?enabled=1240476"
	VerithCorpAnnualURL   = "https://naphsoft.lemonsqueezy.com/checkout/buy/8ec06653-43b4-4dc4-8404-5002993e2058?enabled=1240477"
)

type ProductHandler struct{}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	// Default / Placeholder
	title := "Producto No Encontrado"
	subtitle := ""
	price := ""
	features := []string{}
	images := []string{}
	isDiscount := false
	discountPrice := ""

	if productID == "verith" {
		title = "Verith"
		subtitle = "Toma el control de tu negocio, sin complicaciones."
		price = "$39.99"
		features = []string{
			"Facturación Electrónica SRI (Normativa Actual)",
			"Contabilidad Integral (Ingresos y Egresos)",
			"Acceso completo a todas las funciones",
			"Sincronización automática con el SRI",
			"Actualizaciones normativas constantes",
			"Soporte técnico prioritario",
		}
		images = []string{
			"/static/img/accountableholo/pestana_summary_modo_obscuro.png",
			"/static/img/accountableholo/pestana_summary_modo_claro.png",
			"/static/img/accountableholo/pestana_categorias.png",
			"/static/img/accountableholo/pestana_clientes.png",
			"/static/img/accountableholo/pestana_de_transacciones.png",
			"/static/img/accountableholo/pestana_sri_config_datos_legales.png",
			"/static/img/accountableholo/pestana_sri_config_facturacion.png",
		}
		isDiscount = true
		discountPrice = "$59.99"
	}

	component := pages.ProductDetail(title, subtitle, price, features, images, isDiscount, discountPrice)
	_ = component.Render(r.Context(), w)
}

func (h *ProductHandler) ServePricing(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")
	title := "Verith" // Default

	if productID == "verith" {
		title = "Verith"
	}

	component := pages.Pricing(
		title,
		VerithBasicMonthlyURL,
		VerithBasicAnnualURL,
		VerithPymeMonthlyURL,
		VerithPymeAnnualURL,
		VerithCorpMonthlyURL,
		VerithCorpAnnualURL,
	)
	_ = component.Render(r.Context(), w)
}

func (h *ProductHandler) ServePricingGrid(w http.ResponseWriter, r *http.Request) {
	isAnnual := r.URL.Query().Get("annual") == "true"

	basicLink := VerithBasicMonthlyURL
	pymeLink := VerithPymeMonthlyURL
	corpLink := VerithCorpMonthlyURL

	if isAnnual {
		basicLink = VerithBasicAnnualURL
		pymeLink = VerithPymeAnnualURL
		corpLink = VerithCorpAnnualURL
	}

	sse := datastar.NewSSE(w, r)
	component := products.PricingGrid(isAnnual, basicLink, pymeLink, corpLink)

	// Send the fragment to replace the pricing grid
	if err := sse.PatchElementTempl(component, datastar.WithSelectorID("pricing-grid")); err != nil {
		sse.ConsoleError(err)
	}
}

func (h *ProductHandler) ServeDownload(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")
	title := "Verith"

	if productID == "verith" {
		title = "Verith"
	}

	component := pages.Download(title)
	_ = component.Render(r.Context(), w)
}