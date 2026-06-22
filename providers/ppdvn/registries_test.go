package ppdvn

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tanamoe/urano/test/data/ppdvntest"
)

const (
	PartnerIPM = "Công ty Cổ phần Xuất bản và truyền thông IPM - Số 110, đường Nguyễn Ngọc Nại, Phường Khương Mai, Quận Thanh Xuân, Thành phố Hà Nội"
)

func TestList(t *testing.T) {
	ts := ppdvntest.NewServer()
	t.Cleanup(func() {
		ts.Close()
	})

	tests := []struct {
		name               string
		query              *string
		page               *int
		expectedRegistries []Registry
	}{{
		name:  "SinglePageMultipleRegistries",
		query: new("CLAMP"),
		expectedRegistries: []Registry{{
			ISBN:           "978-604-2-33297-2",
			Title:          "Đội thám tử học viện Clamp Tập 3",
			Author:         "CLAMP",
			Translator:     "Blahira",
			PrintAmount:    200000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "2640-2023/CXBIPH/6-229/KĐ",
		}, {
			ISBN:           "978-604-2-33296-5",
			Title:          "Đội thám tử học viện Clamp Tập 2",
			Author:         "CLAMP",
			Translator:     "Blahira",
			PrintAmount:    200000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "2640-2023/CXBIPH/5-229/KĐ",
		}, {
			ISBN:           "978-604-2-33295-8",
			Title:          "Đội thám tử học viện Clamp Tập 1",
			Author:         "CLAMP",
			Translator:     "Blahira",
			PrintAmount:    200000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "2640-2023/CXBIPH/4-229/KĐ",
		}},
	}, {
		name:               "NoRegistries",
		query:              new("random"),
		expectedRegistries: []Registry{},
	}, {
		name:  "MultiplePageMultipleRegistries",
		query: new("Card Captor Sakura"),
		page:  new(1),
		expectedRegistries: []Registry{{
			ISBN:           "978-604-2-41584-2",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 16",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/59-44/KĐ",
		}, {
			ISBN:           "978-604-2-41583-5",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 15",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/58-44/KĐ",
		}, {
			ISBN:           "978-604-2-41582-8",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 14",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/57-44/KĐ",
		}, {
			ISBN:           "978-604-2-41581-1",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 13",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/56-44/KĐ",
		}, {
			ISBN:           "978-604-2-41580-4",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 12",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/55-44/KĐ",
		}, {
			ISBN:           "978-604-2-41579-8",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 11",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/54-44/KĐ",
		}, {
			ISBN:           "978-604-2-41578-1",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 10",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/53-44/KĐ",
		}, {
			ISBN:           "978-604-2-41577-4",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 9",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/52-44/KĐ",
		}, {
			ISBN:           "978-604-2-41576-7",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 8",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/51-44/KĐ",
		}, {
			ISBN:           "978-604-2-41575-0",
			Title:          "Card Captor Sakura - Thẻ bài pha lê Tập 7",
			Author:         "Clamp",
			Translator:     "Simirimi",
			PrintAmount:    90000,
			SelfPublished:  true,
			Partner:        "",
			RegistrationID: "412-2026/CXBIPH/50-44/KĐ",
		}},
	}, {
		name:  "MultiplePageMultipleRegistriesWithPartner",
		query: new("Cửu Long"),
		page:  new(1),
		expectedRegistries: []Registry{{
			ISBN:           "978-604-40-4056-1",
			Title:          "Chuyện tình Cửu Long - 3",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "1406-2025/CXBIPH/9-74/DT",
		}, {
			ISBN:           "978-604-40-7925-7",
			Title:          "Chuyện tình Cửu Long - 5",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "591-2025/CXBIPH/20-31/DT",
		}, {
			ISBN:           "978-604-40-7924-0",
			Title:          "Chuyện tình Cửu Long - 10",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "591-2025/CXBIPH/19-31/DT",
		}, {
			ISBN:           "978-604-40-7923-3",
			Title:          "Chuyện tình Cửu Long - 9",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "591-2025/CXBIPH/18-31/DT",
		}, {
			ISBN:           "978-604-40-7922-6",
			Title:          "Chuyện tình Cửu Long - 8",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "591-2025/CXBIPH/17-31/DT",
		}, {
			ISBN:           "978-604-40-4061-5",
			Title:          "Chuyện tình Cửu Long - 8",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "2400-2024/CXBIPH/12-101/DT",
		}, {
			ISBN:           "978-604-40-4060-8",
			Title:          "Chuyện tình Cửu Long - 7",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "2400-2024/CXBIPH/11-101/DT",
		}, {
			ISBN:           "978-604-40-4059-2",
			Title:          "Chuyện tình Cửu Long - 6",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "2400-2024/CXBIPH/10-101/DT",
		}, {
			ISBN:           "978-604-40-4058-5",
			Title:          "Chuyện tình Cửu Long - 5",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "2400-2024/CXBIPH/9-101/DT",
		}, {
			ISBN:           "978-604-40-4057-8",
			Title:          "Chuyện tình Cửu Long - 4",
			Author:         "Jun Mayuzuki",
			Translator:     "Thế Đăng",
			PrintAmount:    10000,
			SelfPublished:  false,
			Partner:        PartnerIPM,
			RegistrationID: "2400-2024/CXBIPH/8-101/DT",
		}},
	}}

	client, err := NewClient(WithHTTPClient(ts.Client()), WithDomain(ts.URL))
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			registries, err := client.List(t.Context(), ListParams{
				Query: test.query,
				Page:  test.page,
			})
			require.NoError(t, err)

			assert.EqualValues(t, test.expectedRegistries, registries)
		})
	}
}
