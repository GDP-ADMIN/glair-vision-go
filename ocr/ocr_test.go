package ocr

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/glair-ai/glair-vision-go"
	"github.com/stretchr/testify/assert"
)

func successServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success"}`))
	}))
}

func newConfig(url string) *glair.Config {
	return glair.NewConfig("u", "p", "k").WithBaseURL(url)
}

func TestNew(t *testing.T) {
	config := glair.NewConfig("", "", "")
	ocr := New(config)

	assert.Equal(t, config, ocr.config)
}

func TestKTP(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	// file read error
	_, err := ocr.KTP(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	// success
	_, err = ocr.KTP(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/ktp.jpeg"})
	assert.NoError(t, err)
}

func TestKTPRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.KTPRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/ktp.jpeg"})
	assert.NoError(t, err)
}

func TestKTPWithQuality(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.KTPWithQuality(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.KTPWithQuality(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/ktp.jpeg"})
	assert.NoError(t, err)
}

func TestKTPWithQualityRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.KTPWithQualityRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/ktp.jpeg"})
	assert.NoError(t, err)
}

func TestNPWP(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.NPWP(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.NPWP(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/npwp.jpg"})
	assert.NoError(t, err)
}

func TestNPWPRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.NPWPRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/npwp.jpg"})
	assert.NoError(t, err)
}

func TestKK(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.KK(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.KK(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/kk.jpg"})
	assert.NoError(t, err)
}

func TestKKRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.KKRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/kk.jpg"})
	assert.NoError(t, err)
}

func TestSTNK(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.STNK(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.STNK(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/stnk.jpg"})
	assert.NoError(t, err)
}

func TestSTNKRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.STNKRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/stnk.jpg"})
	assert.NoError(t, err)
}

func TestSIM(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.SIM(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.SIM(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/sim.jpg"})
	assert.NoError(t, err)
}

func TestSIMRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.SIMRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/sim.jpg"})
	assert.NoError(t, err)
}

func TestBPKB(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.BPKB(context.Background(), glair.BPKBInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	// success without page
	_, err = ocr.BPKB(context.Background(), glair.BPKBInput{Image: "../examples/ocr/images/bpkb.pdf"})
	assert.NoError(t, err)

	// success with page
	_, err = ocr.BPKB(context.Background(), glair.BPKBInput{Image: "../examples/ocr/images/bpkb.pdf", Page: glair.Int(1)})
	assert.NoError(t, err)
}

func TestBPKBRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.BPKBRaw(context.Background(), glair.BPKBInput{Image: "../examples/ocr/images/bpkb.pdf"})
	assert.NoError(t, err)

	_, err = ocr.BPKBRaw(context.Background(), glair.BPKBInput{Image: "../examples/ocr/images/bpkb.pdf", Page: glair.Int(1)})
	assert.NoError(t, err)
}

func TestPassport(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.Passport(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.Passport(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/passport.jpeg"})
	assert.NoError(t, err)
}

func TestPassportRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.PassportRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/passport.jpeg"})
	assert.NoError(t, err)
}

func TestPlate(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.Plate(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.Plate(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/plate.jpg"})
	assert.NoError(t, err)
}

func TestPlateRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.PlateRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/plate.jpg"})
	assert.NoError(t, err)
}

func TestGeneralDocument(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.GeneralDocument(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.GeneralDocument(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/general-document.jpg"})
	assert.NoError(t, err)
}

func TestGeneralDocumentRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.GeneralDocumentRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/general-document.jpg"})
	assert.NoError(t, err)
}

func TestInvoice(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.Invoice(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.Invoice(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/invoice.jpg"})
	assert.NoError(t, err)
}

func TestInvoiceRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.InvoiceRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/invoice.jpg"})
	assert.NoError(t, err)
}

func TestReceipt(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.Receipt(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.Receipt(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/receipt.jpg"})
	assert.NoError(t, err)
}

func TestReceiptRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.ReceiptRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/receipt.jpg"})
	assert.NoError(t, err)
}

func TestBankStatement(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.BankStatement(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.BankStatement(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/bank-statement.jpg"})
	assert.NoError(t, err)
}

func TestBankStatementRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.BankStatementRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/bank-statement.jpg"})
	assert.NoError(t, err)
}

func TestSKPR(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.SKPR(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"})
	assert.Error(t, err)

	_, err = ocr.SKPR(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/skpr.jpg"})
	assert.NoError(t, err)
}

func TestSKPRRaw(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	_, err := ocr.SKPRRaw(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/skpr.jpg"})
	assert.NoError(t, err)
}

func TestKTPSessions(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	// without cancel URL
	_, err := ocr.KTPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
	})
	assert.NoError(t, err)

	// with cancel URL
	_, err = ocr.KTPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
		CancelURL:  glair.String("https://example.com/cancel"),
	})
	assert.NoError(t, err)
}

func TestNPWPSessions(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	// without cancel URL
	_, err := ocr.NPWPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
	})
	assert.NoError(t, err)

	// with cancel URL
	_, err = ocr.NPWPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
		CancelURL:  glair.String("https://example.com/cancel"),
	})
	assert.NoError(t, err)
}
