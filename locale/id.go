package locale

var ID map[Message]string = map[Message]string{
	MessageWelcome:                "Halo, {{user}}!\n\nSelamat datang di {{group}}. Pastikan kamu baca pinned message ya.",
	MessageKick:                   "{{user}} telah di kick karena tidak menyelesaikan captcha.",
	MessageJoin:                   "Halo, {{user}}!\n\nSebelum melanjutkan, selesaikan captcha ini dulu. Kamu punya waktu 1 menit dari sekarang.\n\n{{captcha}}",
	MessageWrongAnswerLettersOnly: "Jawaban captcha salah. Hanya huruf saja yang diperbolehkan. Kamu punya {{remaining}} detik lagi untuk menyelesaikan.",
	MessageWrongAnswer:            "Jawaban captcha salah, harap coba lagi. Kamu punya {{remaining}} detik lagi untuk menyelesaikan.",
	MessageNonText:                "Hai, {{user}}. Selesaikan captcha terlebih dahulu ya. Kamu punya waktu {{remaining}} detik lagi.",
}
