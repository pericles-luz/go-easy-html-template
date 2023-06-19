package easy_html_template_test

import (
	"testing"

	"github.com/pericles-luz/go-easy-html-template/pkg/easy_html_template"
	"github.com/stretchr/testify/require"
)

func TestTemplate_TranslateText(t *testing.T) {
	template, err := easy_html_template.NewTemplate()
	require.NoError(t, err)
	template.SetText("testando se {{$nome}} aparece junto com {{$sobrenome}} na idade de {{$idade}} anos")
	template.SetData(map[string]string{
		"nome":      "Pericles",
		"sobrenome": "Luz",
		"idade":     "46",
	})
	result, err := template.GetTranslated()
	require.NoError(t, err)
	require.Equal(t, "testando se Pericles aparece junto com Luz na idade de 46 anos", result)
}

func TestTemplate_TranslateTextWithMissingData(t *testing.T) {
	template, err := easy_html_template.NewTemplate()
	require.NoError(t, err)
	template.SetText("testando se {{$nome}} aparece junto com {{$sobrenome}} na idade de {{$idade}} anos")
	template.SetData(map[string]string{
		"nome":      "Pericles",
		"sobrenome": "Luz",
	})
	_, err = template.GetTranslated()
	require.NotNil(t, err)
}

func TestTemplate_TextWithDataMap(t *testing.T) {
	template, err := easy_html_template.NewTemplate()
	require.NoError(t, err)
	template.SetText("testando se {{.nome}} aparece junto com {{.sobrenome}} na idade de {{.idade}} anos")
	template.SetData(map[string]string{
		"nome":      "Pericles",
		"sobrenome": "Luz",
		"idade":     "46",
	})
	result, err := template.GetTranslated()
	require.NoError(t, err)
	require.Equal(t, "testando se Pericles aparece junto com Luz na idade de 46 anos", result)
}

func TestTemplate_TextWithDataMapAndData(t *testing.T) {
	template, err := easy_html_template.NewTemplate()
	require.NoError(t, err)
	template.SetText("testando se {{$nome}} aparece junto com {{.sobrenome}} na idade de {{$idade}} anos")
	template.SetData(map[string]string{
		"nome":      "Pericles",
		"sobrenome": "Luz",
		"idade":     "46",
	})
	result, err := template.GetTranslated()
	require.NoError(t, err)
	require.Equal(t, "testando se Pericles aparece junto com Luz na idade de 46 anos", result)
}
