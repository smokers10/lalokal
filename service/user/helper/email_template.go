package helper

import "fmt"

func ForgotPasswordEmailTemplate(requester_name string, otp string, link string) string {
	return fmt.Sprintf(`
	<!doctype html>
	<html lang="en">
	  <head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Bootstrap demo</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-0evHe/X+R7YkIZDRvuzKMRqM+OrBnVFBL6DOitfPri4tjfHxaWutUpFmBp4vmVor" crossorigin="anonymous">
	  </head>
	  <body>
		<div class="container">
			<center>
				<h2>Lalokal!</h2>
				<h3>Atur Ulang Password</h3>
			</center>
			<hr>
			<div class="mx-4">
				<p>
					<b>Hallo %s,</b><br>
					Anda lupa password? <br>
					Kami menerima pengajuan atur ulang password untuk akun Anda.
				</p>
				<p>
					Silahkan tekan tombol dibawah untuk melakukan atur ulang password dan masukan kode rahasia dibawah ini: <br>
					Kode Rahasia : <h3>%s</h3>
					<br><br>
					<a href="%s" class="btn btn-primary">
						Atur Ulang Password
					</a>
					<br><br>
					atau copy dan paste link dibawah ini: <br>
					%s
				</p>
				<p>Abaikan pesan ini jika Anda merasa tidak melakukan aksi ini!</p>
				<p>
					Salam Kami,<br>
					Lalokal
				</p>
			</div>
			<hr>
		</div>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/js/bootstrap.bundle.min.js" integrity="sha384-pprn3073KE6tl6bjs2QrFaJGz5/SUsLqktiwsUTF55Jfv3qYSDhgCecCxMW52nD2" crossorigin="anonymous"></script>
	  </body>
	</html>`, requester_name, otp, link, link)
}

func VerificationEmailTemplate(otp string) string {
	return fmt.Sprintf(`
	<!doctype html>
	<html lang="en">
	  <head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Bootstrap demo</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-0evHe/X+R7YkIZDRvuzKMRqM+OrBnVFBL6DOitfPri4tjfHxaWutUpFmBp4vmVor" crossorigin="anonymous">
	  </head>
	  <body>
		<div class="container">
			<center>
				<h2>Lalokal!</h2>
				<h3>Verifikasi Email</h3>
			</center>
			<hr>
			<div class="mx-4">
				<p>
					<b>Hallo!</b><br>
					Kami senang Anda bergabung dengan lalokal. Untuk memulai silahkan masukan kode verifikasi dibawah ini: <br>
					OTP : <h3>%s</h3>
				</p>
				<p>
					Selamat Datang di Lalokal!<br>
					Lalokal Team
				</p>
			</div>
			<hr>
		</div>
		<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.0-beta1/dist/js/bootstrap.bundle.min.js" integrity="sha384-pprn3073KE6tl6bjs2QrFaJGz5/SUsLqktiwsUTF55Jfv3qYSDhgCecCxMW52nD2" crossorigin="anonymous"></script>
	  </body>
	</html>
	`, otp)
}
