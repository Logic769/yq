package yqlib

import (
	"testing"

	"github.com/mikefarah/yq/v4/test"
)

func TestYamlEncoderPreservesQuotedSingleQuoteScalars(t *testing.T) {
	scenarios := []formatScenario{
		{
			description: "empty single quoted scalar",
			input:       "''",
			expected:    "''\n",
		},
		{
			description: "single quoted scalar containing only quotes",
			input:       "''''",
			expected:    "''''\n",
		},
		{
			description: "single quoted scalar is still unwrapped",
			input:       "'hello'",
			expected:    "hello\n",
		},
	}

	for _, s := range scenarios {
		t.Run(s.description, func(t *testing.T) {
			result := mustProcessFormatScenario(s, NewYamlDecoder(ConfiguredYamlPreferences), NewYamlEncoder(ConfiguredYamlPreferences))
			test.AssertResultWithContext(t, s.expected, result, s.description)
		})
	}
}
