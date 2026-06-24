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

	err := ocr.KTP(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.KTP(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/ktp.jpeg"}, nil)
	assert.NoError(t, err)
}

func TestKTPWithQuality(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.KTPWithQuality(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.KTPWithQuality(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/ktp.jpeg"}, nil)
	assert.NoError(t, err)
}

func TestNPWP(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.NPWP(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.NPWP(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/npwp.jpg"}, nil)
	assert.NoError(t, err)
}

func TestKK(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.KK(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.KK(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/kk.jpg"}, nil)
	assert.NoError(t, err)
}

func TestSTNK(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.STNK(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.STNK(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/stnk.jpg"}, nil)
	assert.NoError(t, err)
}

func TestSIM(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.SIM(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.SIM(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/sim.jpg"}, nil)
	assert.NoError(t, err)
}

func TestBPKB(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.BPKB(context.Background(), glair.BPKBInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	// success without page
	err = ocr.BPKB(context.Background(), glair.BPKBInput{Image: "../examples/ocr/images/bpkb.pdf"}, nil)
	assert.NoError(t, err)

	// success with page
	err = ocr.BPKB(context.Background(), glair.BPKBInput{Image: "../examples/ocr/images/bpkb.pdf", Page: glair.Int(1)}, nil)
	assert.NoError(t, err)
}

func TestPassport(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.Passport(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.Passport(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/passport.jpeg"}, nil)
	assert.NoError(t, err)
}

func TestPlate(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.Plate(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.Plate(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/plate.jpg"}, nil)
	assert.NoError(t, err)
}

func TestGeneralDocument(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.GeneralDocument(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.GeneralDocument(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/general-document.jpg"}, nil)
	assert.NoError(t, err)
}

func TestInvoice(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.Invoice(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.Invoice(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/invoice.jpg"}, nil)
	assert.NoError(t, err)
}

func TestReceipt(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.Receipt(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.Receipt(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/receipt.jpg"}, nil)
	assert.NoError(t, err)
}

func TestBankStatement(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.BankStatement(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.BankStatement(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/bank-statement.jpg"}, nil)
	assert.NoError(t, err)
}

func TestSKPR(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	err := ocr.SKPR(context.Background(), glair.OCRInput{Image: "does-not-exist.jpg"}, nil)
	assert.Error(t, err)

	err = ocr.SKPR(context.Background(), glair.OCRInput{Image: "../examples/ocr/images/skpr.jpg"}, nil)
	assert.NoError(t, err)
}

func TestKTPSessions(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	// without cancel URL
	err := ocr.KTPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
	}, nil)
	assert.NoError(t, err)

	// with cancel URL
	err = ocr.KTPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
		CancelURL:  glair.String("https://example.com/cancel"),
	}, nil)
	assert.NoError(t, err)
}

func TestNPWPSessions(t *testing.T) {
	srv := successServer()
	defer srv.Close()

	ocr := New(newConfig(srv.URL))

	// without cancel URL
	err := ocr.NPWPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
	}, nil)
	assert.NoError(t, err)

	// with cancel URL
	err = ocr.NPWPSessions(context.Background(), glair.SessionsInput{
		SuccessURL: "https://example.com/success",
		CancelURL:  glair.String("https://example.com/cancel"),
	}, nil)
	assert.NoError(t, err)
}
