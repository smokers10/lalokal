package keyword

import "lalokal/domain/keyword"

func storeValidation(i *keyword.Keyword) (message string, isfail bool) {
	if i.Keyword == "" {
		return "kata kunci tidak boleh kosong", true
	}

	if i.TopicId == "" {
		return "id topik tidak boleh kosong", true
	}

	return "", false
}
