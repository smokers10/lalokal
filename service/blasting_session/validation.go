package blasting_session

import "lalokal/domain/blasting_session"

func storeValidation(i *blasting_session.BlastingSession) (message string, isfail bool) {
	if i.Message == "" {
		return "pesan tidak boleh kosong", true
	}

	if i.Title == "" {
		return "judul sesi blasting tidak boleh kosong", true
	}

	if i.TopicId == "" {
		return "id topik tidak boleh kosong", true
	}

	return "", false
}

func updateValidation(i *blasting_session.BlastingSession) (message string, isfail bool) {
	if i.Id == "" {
		return "id tidak boleh kosong", true
	}

	if i.Message == "" {
		return "pesan tidak boleh kosong", true
	}

	if i.Title == "" {
		return "judul sesi blasting tidak boleh kosong", true
	}

	if i.TopicId == "" {
		return "id topik tidak boleh kosong", true
	}

	return "", false
}
