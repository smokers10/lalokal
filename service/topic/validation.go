package topic

import "lalokal/domain/topic"

func storeValidation(i *topic.Topic) (string, bool) {
	if i.Title == "" {
		return "judul topik tidak boleh kosong", true
	}

	if i.Description == "" {
		return "deskripsi topik tidak boleh kosong", true
	}

	if i.UserId == "" {
		return "id topik tidak boleh kosong", true
	}

	return "", false
}

func updateValidation(i *topic.Topic) (string, bool) {
	if i.Title == "" {
		return "judul topik tidak boleh kosong", true
	}

	if i.Description == "" {
		return "deskripsi topik tidak boleh kosong", true
	}

	if i.UserId == "" {
		return "id pengguna tidak boleh kosong", true
	}

	if i.Id == "" {
		return "id topik tidak boleh kosong", true
	}

	return "", false
}
