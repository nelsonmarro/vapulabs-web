package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nelsonmarro/vapulabs-web/templates/components/products"
	"github.com/nelsonmarro/vapulabs-web/templates/pages"
	"github.com/starfederation/datastar-go/datastar"
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
			"Acceso completo a todas las funciones",
			"Sincronización automática con el SRI",
			"Actualizaciones normativas constantes",
			"Soporte técnico prioritario",
		}
		images = []string{
			"/static/img/accountableholo/fulllogo.png",
			"/static/img/accountableholo/swappy-20260106_181602.png",
			"/static/img/accountableholo/swappy-20260106_181623.png",
			"/static/img/accountableholo/swappy-20260106_181639.png",
			"/static/img/accountableholo/swappy-20260106_181700.png",
			"/static/img/accountableholo/swappy-20260106_181716.png",
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

	basicMonthly := "https://naphsoft.lemonsqueezy.com/checkout/buy/cbb20457-f455-4733-86c5-0a8cae5b2cdb?enabled=1240067"
	basicAnnual := "https://naphsoft.lemonsqueezy.com/checkout/buy/05ee4c14-5ca6-425b-b798-ef0f5eb33085?enabled=1240066"
	// PYME
	pymeMonthly := "https://naphsoft.lemonsqueezy.com/checkout/buy/19af2e66-93f1-4966-b0e6-26852fe96cb1?enabled=1240404"
	pymeAnnual := "https://naphsoft.lemonsqueezy.com/checkout/buy/cd05a576-94cb-41d8-b1ca-6595f14e346d?enabled=1240412"
	// CORP
	corpMonthly := "https://naphsoft.lemonsqueezy.com/checkout/buy/c388e89a-9e0d-4e76-8fb7-772535c61f43?enabled=1240419"
	corpAnnual := "https://naphsoft.lemonsqueezy.com/checkout/buy/2585f4e7-e151-4a2f-b20e-5af68e51a0ed?enabled=1240420"

	if productID == "verith" {
		title = "Verith"
	}

	component := pages.Pricing(title, basicMonthly, basicAnnual, pymeMonthly, pymeAnnual, corpMonthly, corpAnnual)
	_ = component.Render(r.Context(), w)
}

func (h *ProductHandler) ServePricingGrid(w http.ResponseWriter, r *http.Request) {
	isAnnual := r.URL.Query().Get("annual") == "true"

	// BASIC
	basicMonthly := "https://naphsoft.lemonsqueezy.com/checkout/buy/cbb20457-f455-4733-86c5-0a8cae5b2cdb?enabled=1240067"
	basicAnnual := "https://naphsoft.lemonsqueezy.com/checkout/buy/05ee4c14-5ca6-425b-b798-ef0f5eb33085?enabled=1240066"
	// PYME
	pymeMonthly := "https://naphsoft.lemonsqueezy.com/checkout/buy/19af2e66-93f1-4966-b0e6-26852fe96cb1?enabled=1240404"
	pymeAnnual := "https://naphsoft.lemonsqueezy.com/checkout/buy/cd05a576-94cb-41d8-b1ca-6595f14e346d?enabled=1240412"
	// CORP
	corpMonthly := "https://naphsoft.lemonsqueezy.com/checkout/buy/c388e89a-9e0d-4e76-8fb7-772535c61f43?enabled=1240419"
	corpAnnual := "https://naphsoft.lemonsqueezy.com/checkout/buy/2585f4e7-e151-4a2f-b20e-5af68e51a0ed?enabled=1240420"

	basicLink := basicMonthly
	pymeLink := pymeMonthly
	corpLink := corpMonthly

	if isAnnual {
		basicLink = basicAnnual
		pymeLink = pymeAnnual
		corpLink = corpAnnual
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
