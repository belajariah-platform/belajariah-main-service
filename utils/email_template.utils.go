package utils

import (
	"belajariah-main-service/model"
	"fmt"
)

func TemplateChangePassword(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
        <tr>
            <td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
                <div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7;";>
                   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr style="margin:0;padding:0">
                                <td style="text-align: center;">
                                    <img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
                                </td>
                                </tr>
                                <tr>
                                    <td>
                                        <p style="font-size:14px;margin: 10px 20px 10px 20px;">
                                            Assalamu’alaikum, <b style="color:#212121">%s</b>,
                                        </p>
                                    </td>
                                </tr>
                                <tr style="margin:0px 20px 0px 20px;padding:0">
                                    <td style="margin:0px 20px 0px 20px;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:100;margin:10px 20px 20px 20px;padding:0;font-size: 14px;">Password belajariahmu berhasil di ubah.</h5>  
                                         <h5 style="line-height:1.4;color:black;font-weight:100;margin:10px 20px 20px 20px;padding:0;font-size: 14px;">Jika kamu punya pertanyaan silahkan hubungin kontak customer service belajariah di bawah.</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
                        <tbody>
                            <tr>
                                <td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
                                    <p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:32px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                        <tbody>
                                            <tr>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0">
                                                                     <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
                                                                    <p style="margin:0">Ikuti Kami</p>
                                                                </td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0;text-align:right">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:24px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
                                        <tbody>
                                            <tr>
                                                <td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
                                                    <p style="margin:0">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table> 
                </div>
            </td>
        </tr>
    </table>`,
		"100%", "100%", email.UserEmail, "100%",
		email.WhatsApp, "100%", "60%", "100%",
		email.GooglePLay, "100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram, "100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplateRegisterSuccess(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
	<tr>
		<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
			<div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
			   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
					<tbody>
						<tr style="margin:0;padding:0">
							<td style="text-align: center;">
								<img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
							</td>
						</tr>
						<tr>
							<td>
								<p style="font-size:14px;margin: 0px 20px 10px 20px;">Assalamu’alaikum, <b style="color:#212121">%s</b>,</p>
							</td>
						</tr>
						<tr style="margin:0;padding:0">
							<td style="margin:0;padding:0">
								<h3 style="line-height:1.4;color:black;font-weight:700;margin:0px 20px 20px 20px;padding:0;font-size: 14px;">Selamat bergabung di Belajariah! Nikmati manfaat dan keunggulan Belajar Al-Qur’an melalui kelas-kelas terbaik dari kami!</h3>  
							</td>
						</tr>
					</tbody>
				</table>
				<table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 0px auto;width: %s;" cellspacing="0" cellpadding="0">
					<tbody>
						<tr style="margin:0;padding:0">
							<td style="text-align: center;">
								<img src="https://www.belajariah.com/email-assets/img/Assets_Illustrasi_Beljariah4.png" width="auto" height="180px" alt="" style="margin:0px auto 0px auto;">
								<h5 style="line-height:1;color:black;font-weight:700;margin:0px auto 30px auto;padding:0;font-size: 14px;">Belajar dengan Ustadz yang berkompeten</h5>  

							</td>
						</tr>
						<tr style="margin:0;padding:0">
							<td style="text-align: center;">
								<img src="https://www.belajariah.com/email-assets/img/Assets_Illustrasi_Beljariah2.png" width="auto" height="180px" alt="" style="margin:0px auto 0px auto;">
								<h5 style="line-height:1;color:black;font-weight:700;margin:0px auto 30px auto;padding:0;font-size: 14px;">Metode mudah dan menyenangkan</h5>  
							</td>
						</tr>
						<tr style="margin:0;padding:0">
							<td style="text-align: center;">
								<img src="https://www.belajariah.com/email-assets/img/Assets_Illustrasi_Beljariah6.png" width="auto" height="180px" alt="" style="margin:0px auto 0px auto;">
								<h5 style="line-height:1;color:black;font-weight:700;margin:0px auto 30px auto;padding:0;font-size: 14px;">Sertifikat dan raport hasil belajar</h5>  
							</td>
						</tr>
					</tbody>
				</table>
				<table style="max-width:%s;border-spacing:0;width:%s;background-color:transparent;margin: 40px 0px 0px 0px;padding:20px" cellspacing="0" cellpadding="0">
					<tbody>
							<tr style="margin:0;padding:0">
								<td style="margin:0;padding:0">
									<h5 style="line-height:1.4;color:black;font-weight:100;margin:0px 20px 20px 20px;padding:0;font-size: 14px;">Kamu bisa belajar Al-Qur’an kapan saja dan dimana saja di sini!</h5>  
								</td>
							</tr>
							<tr style="margin:0;padding:0">
								<td style="margin:0;padding:0">
									<h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 20px 20px 20px;padding:0;font-size: 14px;">Belajariah Solusi Kamu Belajar Al-Qur’an!!!</h5>  
								</td>
							</tr>
					</tbody>
				</table>
				<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
					<tbody>
						<tr>
							<td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
								<p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
							</td>
						</tr>
						<tr>
							<td style="padding:32px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
									<tbody>
										<tr>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
														</tr>
														<tr>
															<td style="padding:8px 0">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
																<p style="margin:0">Ikuti Kami</p>
															</td>
														</tr>
														<tr>
															<td style="padding:8px 0;text-align:right">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
						<tr>
							<td style="padding:24px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
									<tbody>
										<tr>
											<td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
												<p style="margin:0">%s</p>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
					</tbody>
				</table> 
			</div>
		</td>
	</tr>
	</table>`,
		"100%", "100%", email.UserName,
		"50%", "100%", "100%", "100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplateAccountVerification(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
        <tr>
            <td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
                <div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7;";>
                   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr style="margin:0;padding:0">
                                <td style="text-align: center;">
                                    <img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
                                </td>
                                </tr>
                                <tr>
                                    <td>
                                        <p style="font-size:14px;margin: 10px 20px 10px 20px;">
                                            Assalamu’alaikum, <b style="color:#212121">%s</b>,
                                        </p>
                                    </td>
                                </tr>
                                <tr style="margin:0px 20px 0px 20px;padding:0">
                                    <td style="margin:0px 20px 0px 20px;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:100;margin:10px 20px 20px 20px;padding:0;font-size: 14px;">Aktifkan akunmu dengan kode di bawah ini:</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 0px auto;" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td style="width:180px;padding-right:5px">
                                    <div style="background-color: #fa591d;color: #fff;text-align: center;font-weight: 600;padding: 5px 0px 5px 0px;">
                                        <p style="margin: 5px 0px 5px 0px;font-size: 14px;">Kode verifikasi : <br>%s</p>
                                    </div>
                                 </td>  
                            </tr>
                        </tbody>
                    </table>
                    <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                                <tr style="margin:0;padding:0">
                                    <td style="margin:0;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:100;font-size:20px;margin:20px 20px 10px 20px;padding:0;font-size: 14px;">Kami perlu memastikan bahwa email anda benar dan tidak disalahgunakan oleh pihak yang tidak berkepentingan.</h5>  
                                    </td>
                                </tr>
                                <tr style="margin:0;padding:0">
                                    <td style="margin:0;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:100;font-size:20px;margin:0px 0px 10px 0px;padding:0;font-size: 16px;">Klik <a href="belajariah://app/verif">Link ini</a> untuk ke Halaman Verifikasi Belajariah</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
                        <tbody>
                            <tr>
                                <td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
                                    <p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:32px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                        <tbody>
                                            <tr>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0">
                                                                     <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
                                                                    <p style="margin:0">Ikuti Kami</p>
                                                                </td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0;text-align:right">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:24px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
                                        <tbody>
                                            <tr>
                                                <td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
                                                    <p style="margin:0">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table> 
                </div>
            </td>
        </tr>
    </table>`,
		"100%", "100%",
		email.UserName, email.VerificationCode,
		"100%", "100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)

	return bodyTemp
}

func TemplateBeforeClassExpired(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
	<tr>
		<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
			<div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
			   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
					<tbody>
						<tr style="margin:0;padding:0">
							<td style="text-align: center;">
								<img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
							</td>
							</tr>
							<tr>
								<td>
									<p style="font-size:14px;margin: 0px 20px 10px 20px;">
										Assalamu’alaikum, <b style="color:#212121">%s</b>,
									</p>
								</td>
							</tr>
							<tr style="margin:0;padding:0">
								<td style="margin:0;padding:0">
									<h5 style="margin:0px 20px 20px 20px;line-height:1.4;color:black;font-weight:100;padding:0;font-size: 14px;">Kelas Tahsin akan berakhir pada <b>%s</b></h5>
									<h5 style="margin:0px 20px 20px 20px;line-height:1.4;color:black;font-weight:100;padding:0;font-size: 14px;">Spesial buat kamu ! Kamu dapat memperpanjang masa kelas menggunakan <b>Promo Langganan diskon %s.</b></h5>
									<h5 style="margin:0px 20px 20px 20px;line-height:1.4;color:#fa591d;font-weight:700;padding:0;font-size: 14px;">%d hari lagi Kelasmu akan berakhir !</h5>
									<h5 style="margin:0px 20px 20px 20px;line-height:1.4;color:black;font-weight:100;padding:0;font-size: 14px;">Dapatkan sekarang sebelum kehabisan !</h5>    
								</td>
							</tr>
					</tbody>
				</table>
				<table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 0px auto;width: %s;" cellspacing="0" cellpadding="0">
					<tbody>
						<tr>
							<td style="width:%s;padding-right:5px">
								<a href="#" style="background: linear-gradient(#542f91, #3f1c78, #835bc2);color: #fff;display:block;border: 2.5px solid #552b9c;border-radius:20px;padding:10px 10px;text-align:center;font-weight:bold;text-decoration:none;font-size:14px">Klaim Voucher</a> 
							 </td>  
						</tr>
					</tbody>
				</table>
				<table cellspacing="0" cellpadding="0" style="width:%s">
					<tbody>
						<tr>
							<td style="padding-bottom:20px">
								<table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
									<tbody>

											<tr>
												<td>
													<table cellspacing="0" cellpadding="0" style="width:%s">
														<tbody>
															<td style="padding:25px 20px 16px 0px">
																<h2 style="font-size:18px;font-weight:600;margin:0"></h2>
															</td>
															<tr>
																<td style="vertical-align:top;padding-bottom:10px">
																	<table cellspacing="0" cellpadding="0">
																		<tbody>
																			<tr style="vertical-align:top;">
																				<td>
																					<a href="#" style="text-decoration:none">
																						<img src="https://www.belajariah.com/img-assets/bannerperpanjangkelas.png" style="width:%s;margin: 0px 0px 20px 0px;">
																					</a>
																				</td>
																			</tr>
																		</tbody>
																	</table>  
																</td>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr> 
					</tbody>
				</table>
				<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
					<tbody>
						<tr>
							<td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
								<p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
							</td>
						</tr>
						<tr>
							<td style="padding:32px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
									<tbody>
										<tr>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
														</tr>
														<tr>
															<td style="padding:8px 0">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
																<p style="margin:0">Ikuti Kami</p>
															</td>
														</tr>
														<tr>
															<td style="padding:8px 0;text-align:right">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
						<tr>
							<td style="padding:24px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
									<tbody>
										<tr>
											<td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
												<p style="margin:0">%s</p>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
					</tbody>
				</table> 
			</div>
		</td>
	</tr>
	</table>`,
		"100%", "100%", email.UserName,
		CurrentDateStringCustom(email.ExpiredDate), email.PromoDiscount, email.Count,
		"30%", "30%", "100%", "100%", "100%", "100%", "100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplateClassHasBeenExpired(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
	<tr>
		<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
			<div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
			   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
					<tbody>
						<tr style="margin:0;padding:0">
							<td style="text-align: center;">
								<img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
							</td>
							</tr>
							<tr>
								<td>
									<p style="font-size:14px;margin: 0px 20px 10px 20px;">
										Assalamu’alaikum, <b style="color:#212121">%s</b>,
									</p>
								</td>
							</tr>
							<tr style="margin:0;padding:0">
								<td style="margin:0;padding:0">
									<h5 style="margin:0px 20px 20px 20px;line-height:1.4;color:#fa591d;font-weight:700;padding:0;font-size: 14px;">Kelasmu telah berakhir !</h5>
									<h5 style="margin:0px 20px 20px 20px;line-height:1.4;color:black;font-weight:100;padding:0;font-size: 14px;">Spesial buat kamu ! Kamu dapat memperpanjang masa kelas menggunakan <b>Promo Langganan diskon %s</b></h5>
									<h5 style="margin:0px 20px 20px 20px;line-height:1.4;color:black;font-weight:100;padding:0;font-size: 14px;">Dapatkan sekarang sebelum kehabisan !</h5> 
								</td>
							</tr>
					</tbody>
				</table>
				<table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 0px auto;width:%s;" cellspacing="0" cellpadding="0">
					<tbody>
						<tr>
							<td style="width:%s;padding-right:5px">
								<a href="#" style="background: linear-gradient(#542f91, #3f1c78, #835bc2);color: #fff;display:block;border: 2.5px solid #552b9c;border-radius:20px;padding:10px 10px;text-align:center;font-weight:bold;text-decoration:none;font-size:14px">Klaim Voucher</a> 
							 </td>  
						</tr>
					</tbody>
				</table>
				<table cellspacing="0" cellpadding="0" style="width:%s">
					<tbody>
						<tr>
							<td style="padding-bottom:20px">
								<table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
									<tbody>

											<tr>
												<td>
													<table cellspacing="0" cellpadding="0" style="width:%s">
														<tbody>
															<td style="padding:25px 20px 16px 0px">
																<h2 style="font-size:18px;font-weight:600;margin:0"></h2>
															</td>
															<tr>
																<td style="vertical-align:top;padding-bottom:10px">
																	<table cellspacing="0" cellpadding="0">
																		<tbody>
																			<tr style="vertical-align:top;">
																				<td>
																					<a href="#" style="text-decoration:none">
																						<img src="https://www.belajariah.com/img-assets/bannerperpanjangkelas.png" style="width:%s;margin: 0px 0px 20px 0px;">
																					</a>
																				</td>
																			</tr>
																		</tbody>
																	</table>  
																</td>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr> 
					</tbody>
				</table>
				<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
					<tbody>
						<tr>
							<td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
								<p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
							</td>
						</tr>
						<tr>
							<td style="padding:32px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
									<tbody>
										<tr>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
														</tr>
														<tr>
															<td style="padding:8px 0">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
																<p style="margin:0">Ikuti Kami</p>
															</td>
														</tr>
														<tr>
															<td style="padding:8px 0;text-align:right">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
						<tr>
							<td style="padding:24px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
									<tbody>
										<tr>
											<td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
												<p style="margin:0">%s</p>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
					</tbody>
				</table> 
			</div>
		</td>
	</tr>
	</table>`,
		"100%", "100%", email.UserName, email.PromoDiscount,
		"30%", "30%", "100%", "100%", "100%", "100%", "100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplateWaitingPayment(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
    <table width="%s">
        <tr>
            <td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
                <div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
                   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr style="margin:0;padding:0">
                                <td style="text-align: center;">
                                    <img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
                                </td>
                                </tr>
                                <tr>
                                    <td>
                                        <p style="font-size:14px">
                                            Assalamu’alaikum, <b style="color:#212121">%s</b>,
                                        </p>
                                    </td>
                                </tr>
                                <tr style="margin:0;padding:0">
                                    <td style="margin:0;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 0px 10px 0px;padding:0;font-size: 14px;">Selangkah lagi kamu bisa memulai kelas. Selesaikan sekarang sebelum kelas penuh !
									</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 30px auto;width: %s;" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td style="width:%s;padding-right:5px">
                                    <a href="#" style="background: linear-gradient(#542f91, #3f1c78, #835bc2);color: #fff;display:block;border: 2.5px solid #552b9c;border-radius:20px;padding:10px 10px;text-align:center;font-weight:600;text-decoration:none;font-size:14px">Bayar Sekarang</a> 
                                 </td>  
                            </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;width:%s;margin-bottom:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td>
                                    <table style="background:#f3f4f5;border-radius:8px;padding:20px;width:%s;" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr style="vertical-align:top;padding-bottom:10px">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:black;font-size:14px;margin:0;font-weight: bold;">Rincian Bank</p>
                                                </td>
                                                
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">Bank</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0 0 8px 0">
                                                        <span style="color:rgba(49,53,59,0.96);font-weight:bold;font-size:14px">
                                                            %s a.n %s
                                                        </span>
                                                    </p>
                                                </td>
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">No. Rekening</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0;font-weight:bold;font-size:14px;color:rgba(49,53,59,0.96)">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin:0;width:%s" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td colspan="2">
                                    <h2 style="font-size:16px;font-weight:bold;color:rgba(49,53,59,0.96);margin:0;margin-bottom:10px;font-size:14px">Ringkasan Pembayaran</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">Harga kelas %s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">%s</td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">%s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">- %s</td>
                            </tr>
                            <tr>
                                <td colspan="2">
                                    <span style="display:block;width:%s;height:1px;padding:0;background:#e5e7e9;margin:10px 0"></span>
                                </td>
                            </tr>
                            <tr>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px">Total Bayar</td>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px;text-align:right;color:#fa591d">%s</td>
                            </tr>   
                        </tbody>
                    </table>
                    <table cellspacing="0" cellpadding="0" style="width:%s">
                        <tbody>
                            <tr>
                                <td style="padding:25px 20px 16px 20px">
                                    <h2 style="font-size:14px;font-weight:600;margin:0">Rincian Pesanan Kelas</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding-bottom:20px">
                                    <table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr>
                                                <td colspan="2">
                                                    <p style="margin:0 0 10px 0;">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">No. Invoice:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td colspan="2">
                                                        <p style="margin:0 0 25px 0">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">Metode Pembayaran:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <table cellspacing="0" cellpadding="0" style="width:%s">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="text-align: center;">
                                                                        <img src="%s" style="width:auto;height:180px;">
                                                                    </td>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr> 
                        </tbody>
                    </table>
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
                        <tbody>
                            <tr>
                                <td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
                                    <p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:32px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                        <tbody>
                                            <tr>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
                                                                    <p style="margin:0">Ikuti Kami</p>
                                                                </td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0;text-align:right">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:24px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
                                        <tbody>
                                            <tr>
                                                <td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
                                                    <p style="margin:0">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table> 
                </div>
            </td>
        </tr>
    </table> `,
		"100%", "100%",
		email.UserName,
		"30%", "30%", "100%",
		"100%", "30%", "30%", "70%",
		email.PaymentMethod,
		email.AccountName,
		"30%", "70%",
		email.AccountNumber,
		"100%",
		email.ClassName,
		FormatRupiah(email.ClassPrice),
		email.PromoDiscount,
		FormatRupiah(email.PromoPrice),
		"100%",
		FormatRupiah(email.TotalTransfer),
		"100%", "100%",
		email.InvoiceNumber,
		email.PaymentMethod,
		"100%",
		email.ClassImage,
		"100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplatePayment2HoursBeforeExpired(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
        <tr>
            <td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
                <div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
                   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr style="margin:0;padding:0">
                                <td style="text-align: center;">
                                    <img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
                                </td>
                                </tr>
                                <tr>
                                    <td>
                                        <p style="font-size:14px">
                                            Assalamu’alaikum, <b style="color:#212121">%s</b>,
                                        </p>
                                    </td>
                                </tr>
                                <tr style="margin:0;padding:0">
                                    <td style="margin:0;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 0px 10px 0px;padding:0;font-size: 14px;">Pesanan Anda masih menunggu pembayaran. Mohon lakukan pembayaran sebelum pukul %s pada %s atau pesanan Anda akan dibatalkan secara otomatis.
									</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 30px auto;width: %s;" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td style="width:%s;padding-right:5px">
                                    <a href="#" style="background: linear-gradient(#542f91, #3f1c78, #835bc2);color: #fff;display:block;border: 2.5px solid #552b9c;border-radius:20px;padding:10px 10px;text-align:center;font-weight:600;text-decoration:none;font-size:14px">Bayar Sekarang</a> 
                                 </td>  
                            </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;width:%s;margin-bottom:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td>
                                    <table style="background:#f3f4f5;border-radius:8px;padding:20px;width:%s;" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr style="vertical-align:top;padding-bottom:10px">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:black;font-size:14px;margin:0;font-weight: bold;">Rincian Bank</p>
                                                </td>
                                                
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">Bank</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0 0 8px 0">
                                                        <span style="color:rgba(49,53,59,0.96);font-weight:bold;font-size:14px">
                                                            %s a.n %s
                                                        </span>
                                                    </p>
                                                </td>
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">No. Rekening</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0;font-weight:bold;font-size:14px;color:rgba(49,53,59,0.96)">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin:0;width:%s" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td colspan="2">
                                    <h2 style="font-size:16px;font-weight:bold;color:rgba(49,53,59,0.96);margin:0;margin-bottom:10px;font-size:14px">Ringkasan Pembayaran</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">Harga kelas %s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">%s</td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">%s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">- %s</td>
                            </tr>
                            <tr>
                                <td colspan="2">
                                    <span style="display:block;width:%s;height:1px;padding:0;background:#e5e7e9;margin:10px 0"></span>
                                </td>
                            </tr>
                            <tr>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px">Total Bayar</td>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px;text-align:right;color:#fa591d">%s</td>
                            </tr>   
                        </tbody>
                    </table>
                    <table cellspacing="0" cellpadding="0" style="width:%s">
                        <tbody>
                            <tr>
                                <td style="padding:25px 20px 16px 20px">
                                    <h2 style="font-size:14px;font-weight:600;margin:0">Rincian Pesanan Kelas</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding-bottom:20px">
                                    <table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr>
                                                <td colspan="2">
                                                    <p style="margin:0 0 10px 0;">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">No. Invoice:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td colspan="2">
                                                        <p style="margin:0 0 25px 0">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">Metode Pembayaran:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <table cellspacing="0" cellpadding="0" style="width:%s">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="text-align: center;">
                                                                        <img src="%s" style="width:auto;height:180px;">
                                                                    </td>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr> 
                        </tbody>
                    </table>
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
                        <tbody>
                            <tr>
                                <td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
                                    <p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:32px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                        <tbody>
                                            <tr>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
                                                                    <p style="margin:0">Ikuti Kami</p>
                                                                </td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0;text-align:right">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:24px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
                                        <tbody>
                                            <tr>
                                                <td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
                                                    <p style="margin:0">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table> 
                </div>
            </td>
        </tr>
    </table> `,
		"100%", "100%",
		email.UserName,
		CurrentTime(email.ExpiredDate),
		CurrentDate(email.ExpiredDate),
		"30%", "30%", "100%",
		"100%", "30%", "30%", "70%",
		email.PaymentMethod,
		email.AccountName,
		"30%", "70%",
		email.AccountNumber,
		"100%",
		email.ClassName,
		FormatRupiah(email.ClassPrice),
		email.PromoDiscount,
		FormatRupiah(email.PromoPrice),
		"100%",
		FormatRupiah(email.TotalTransfer),
		"100%", "100%",
		email.InvoiceNumber,
		email.PaymentMethod,
		"100%",
		email.ClassImage,
		"100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplateUploadPayment(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
        <tr>
            <td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
                <div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
                   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr style="margin:0;padding:0">
                                <td style="text-align: center;">
                                    <img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
                                </td>
                                </tr>
                                <tr>
                                    <td>
                                        <p style="font-size:14px">
                                            Assalamu’alaikum, <b style="color:#212121">%s</b>,
                                        </p>
                                    </td>
                                </tr>
                                <tr style="margin:0;padding:0">
                                    <td style="margin:0;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 0px 10px 0px;padding:0;font-size: 14px;">Bukti pembayaran kelas %s dengan nomor pesanan %s BERHASIL DIKIRIM ! Silahkan menunggu konfirmasi ..
									</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                   
                    <table style="padding:0 20px;width:%s;margin-bottom:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td>
                                    <table style="background:#f3f4f5;border-radius:8px;padding:20px;width:%s;" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr style="vertical-align:top;padding-bottom:10px">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:black;font-size:14px;margin:0;font-weight: bold;">Rincian Bank</p>
                                                </td>
                                                
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">Bank</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0 0 8px 0">
                                                        <span style="color:rgba(49,53,59,0.96);font-weight:bold;font-size:14px">
                                                            %s a.n %s
                                                        </span>
                                                    </p>
                                                </td>
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">No. Rekening</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0;font-weight:bold;font-size:14px;color:rgba(49,53,59,0.96)">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin:0;width:%s" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td colspan="2">
                                    <h2 style="font-size:16px;font-weight:bold;color:rgba(49,53,59,0.96);margin:0;margin-bottom:10px;font-size:14px">Ringkasan Pembayaran</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">Harga kelas %s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">%s</td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">%s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">- %s</td>
                            </tr>
                            <tr>
                                <td colspan="2">
                                    <span style="display:block;width:%s;height:1px;padding:0;background:#e5e7e9;margin:10px 0"></span>
                                </td>
                            </tr>
                            <tr>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px">Total Bayar</td>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px;text-align:right;color:#fa591d">Rp%s</td>
                            </tr>   
                        </tbody>
                    </table>
                    <table cellspacing="0" cellpadding="0" style="width:%s">
                        <tbody>
                            <tr>
                                <td style="padding:25px 20px 16px 20px">
                                    <h2 style="font-size:14px;font-weight:600;margin:0">Rincian Pesanan Kelas</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding-bottom:20px">
                                    <table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr>
                                                <td colspan="2">
                                                    <p style="margin:0 0 10px 0;">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">No. Invoice:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td colspan="2">
                                                        <p style="margin:0 0 25px 0">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">Metode Pembayaran:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <table cellspacing="0" cellpadding="0" style="width:%s">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="text-align: center;">
                                                                        <img src="%s" style="width:auto;height:180px;">
                                                                    </td>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr> 
                        </tbody>
                    </table>
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
                        <tbody>
                            <tr>
                                <td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
                                    <p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:32px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                        <tbody>
                                            <tr>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
                                                                    <p style="margin:0">Ikuti Kami</p>
                                                                </td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0;text-align:right">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:24px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
                                        <tbody>
                                            <tr>
                                                <td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
                                                    <p style="margin:0">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table> 
                </div>
            </td>
        </tr>
    </table> `,
		"100%", "100%",
		email.UserName,
		email.ClassName,
		email.InvoiceNumber,
		"100%", "100%", "30%", "30%", "70%",
		email.PaymentMethod,
		email.AccountName,
		"30%", "70%",
		email.AccountNumber,
		"100%",
		email.ClassName,
		FormatRupiah(email.ClassPrice),
		email.PromoDiscount,
		FormatRupiah(email.PromoPrice),
		"100%",
		FormatRupiah(email.TotalTransfer),
		"100%", "100%",
		email.InvoiceNumber,
		email.PaymentMethod,
		"100%",
		email.ClassImage,
		"100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplatePaymentSuccess(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
        <tr>
            <td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
                <div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
                   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr style="margin:0;padding:0">
                                <td style="text-align: center;">
                                    <img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
                                </td>
                                </tr>
                                <tr>
                                    <td>
                                        <p style="font-size:14px">
                                            Assalamu’alaikum, <b style="color:#212121">%s</b>,
                                        </p>
                                    </td>
                                </tr>
                                <tr style="margin:0;padding:0">
                                    <td style="margin:0;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 0px 10px 0px;padding:0;font-size: 14px;">Pembelian Kelas %s dengan nomor pesanan %s BERHASIL DIKONFIRMASI ! Kelas siap dimulai ..
									</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 30px auto;width: %s;" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td style="width:%s;padding-right:5px">
                                    <a href="#" style="background: linear-gradient(#542f91, #3f1c78, #835bc2);color: #fff;display:block;border: 2.5px solid #552b9c;border-radius:20px;padding:10px 10px;text-align:center;font-weight:600;text-decoration:none;font-size:14px">Bayar Sekarang</a> 
                                 </td>  
                            </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;width:%s;margin-bottom:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td>
                                    <table style="background:#f3f4f5;border-radius:8px;padding:20px;width:%s;" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr style="vertical-align:top;padding-bottom:10px">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:black;font-size:14px;margin:0;font-weight: bold;">Rincian Bank</p>
                                                </td>
                                                
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s" style="padding:0 0 15px 0">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">Bank</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0 0 8px 0">
                                                        <span style="color:rgba(49,53,59,0.96);font-weight:bold;font-size:14px">
                                                            %s a.n %s
                                                        </span>
                                                    </p>
                                                </td>
                                            </tr>
                                            <tr style="vertical-align:top">
                                                <td width="%s">
                                                    <p style="color:rgba(49,53,59,0.68);font-size:12px;margin:0">No. Rekening</p>
                                                </td>
                                                <td width="%s">
                                                    <p style="margin:0;font-weight:bold;font-size:14px;color:rgba(49,53,59,0.96)">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin:0;width:%s" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td colspan="2">
                                    <h2 style="font-size:16px;font-weight:bold;color:rgba(49,53,59,0.96);margin:0;margin-bottom:10px;font-size:14px">Ringkasan Pembayaran</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">Harga kelas %s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">%s</td>
                            </tr>
                            <tr>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;margin-bottom:8px">%s</td>
                                <td style="color:rgba(49,53,59,0.68);font-size:14px;padding:6px 0;text-align:right;font-weight:bold;color:rgba(49,53,59,0.96)">- %s</td>
                            </tr>
                            <tr>
                                <td colspan="2">
                                    <span style="display:block;width:%s;height:1px;padding:0;background:#e5e7e9;margin:10px 0"></span>
                                </td>
                            </tr>
                            <tr>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px">Total Bayar</td>
                                <td style="font-size:14px;font-weight:bold;padding-bottom:16px;text-align:right;color:#fa591d">Rp%s</td>
                            </tr>   
                        </tbody>
                    </table>
                    <table cellspacing="0" cellpadding="0" style="width:%s">
                        <tbody>
                            <tr>
                                <td style="padding:25px 20px 16px 20px">
                                    <h2 style="font-size:14px;font-weight:600;margin:0">Rincian Pesanan Kelas</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding-bottom:20px">
                                    <table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr>
                                                <td colspan="2">
                                                    <p style="margin:0 0 10px 0;">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">No. Invoice:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td colspan="2">
                                                        <p style="margin:0 0 25px 0">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">Metode Pembayaran:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <table cellspacing="0" cellpadding="0" style="width:%s">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="text-align: center;">
                                                                        <img src="%s" style="width:auto;height:180px;">
                                                                    </td>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr> 
                        </tbody>
                    </table>
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
                        <tbody>
                            <tr>
                                <td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
                                    <p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:32px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                        <tbody>
                                            <tr>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
                                                                    <p style="margin:0">Ikuti Kami</p>
                                                                </td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0;text-align:right">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:24px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
                                        <tbody>
                                            <tr>
                                                <td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
                                                    <p style="margin:0">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table> 
                </div>
            </td>
        </tr>
    </table> `,
		"100%", "100%",
		email.UserName,
		email.ClassName,
		email.InvoiceNumber,
		"30%", "30%", "100%",
		"100%", "30%", "30%", "70%",
		email.PaymentMethod,
		email.AccountName,
		"30%", "70%",
		email.AccountNumber,
		"100%",
		email.ClassName,
		FormatRupiah(email.ClassPrice),
		email.PromoDiscount,
		FormatRupiah(email.PromoPrice),
		"100%",
		FormatRupiah(email.TotalTransfer),
		"100%", "100%",
		email.InvoiceNumber,
		email.PaymentMethod,
		"100%",
		email.ClassImage,
		"100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplatePaymentFailed(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
    <table width="%s">
        <tr>
            <td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
                <div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
                   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr style="margin:0;padding:0">
                                <td style="text-align: center;">
                                    <img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
                                </td>
                                </tr>
                                <tr>
                                    <td>
                                        <p style="font-size:14px;margin-bottom:10px">
                                            Assalamu’alaikum, <b style="color:#212121">%s</b>,
                                        </p>
                                    </td>
                                </tr>
                                <tr style="margin:0;padding:0">
                                    <td style="margin:0;padding:0">
                                        <h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 0px 10px 0px;padding:0;font-size: 14px;">Mohon maaf pembelian kelas telah dibatalkan, karena anda tidak menyelesaikan pembayaran </h5>  
 										<h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 0px 10px 0px;padding:0;font-size: 14px;">Silahkan melakukan pemesanan kelas kembali di Aplikasi Belajariah</h5>  
                                    </td>
                                </tr>
                        </tbody>
                    </table>
                    <table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 30px auto;width: %s;" cellspacing="0" cellpadding="0">
                        <tbody>
                            <tr>
                                <td>
                                    <a href="#" style="background: linear-gradient(#542f91, #3f1c78, #835bc2);color: #fff;display:block;border: 2.5px solid #552b9c;border-radius:20px;padding:10px 10px;text-align:center;font-weight:600;text-decoration:none;font-size:14px">Belajar Al-Quran Sekarang</a> 
                                 </td>  
                            </tr>
                        </tbody>
                    </table>
                    <table cellspacing="0" cellpadding="0" style="width:%s">
                        <tbody>
                            <tr>
                                <td style="padding:25px 20px 16px 20px">
                                    <h2 style="font-size:14px;font-weight:600;margin:0">Rincian Pesanan Kelas</h2>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding-bottom:20px">
                                    <table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
                                        <tbody>
                                            <tr>
                                                <td colspan="2">
                                                    <p style="margin:0 0 10px 0;">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">No. Invoice:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td colspan="2">
                                                        <p style="margin:0 0 25px 0">
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">Kelas:</span>
                                                            <span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s
                                                        </p>
                                                    </td>
                                                </tr>
                                                <tr>
                                                    <td>
                                                        <table cellspacing="0" cellpadding="0" style="width:%s">
                                                            <tbody>
                                                                <tr>
                                                                    <td style="text-align: center;">
                                                                        <img src="%s" style="width:auto;height:180px;">
                                                                    </td>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr> 
                        </tbody>
                    </table>
                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
                        <tbody>
                            <tr>
                                <td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
                                    <p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:32px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                        <tbody>
                                            <tr>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                                <td valign="top" width="%s">
                                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
                                                        <tbody>
                                                            <tr>
                                                                <td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
                                                                    <p style="margin:0">Ikuti Kami</p>
                                                                </td>
                                                            </tr>
                                                            <tr>
                                                                <td style="padding:8px 0;text-align:right">
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                    <a href="%s" style="text-decoration:none">
                                                                        <img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
                                                                    </a>
                                                                </td>
                                                            </tr>
                                                        </tbody>
                                                    </table>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                            <tr>
                                <td style="padding:24px 20px 0;background:#ffffff">
                                    <table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
                                        <tbody>
                                            <tr>
                                                <td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
                                                    <p style="margin:0">%s</p>
                                                </td>
                                            </tr>
                                        </tbody>
                                    </table>
                                </td>
                            </tr>
                        </tbody>
                    </table> 
                </div>
            </td>
        </tr>
    </table>`,
		"100%", "100%", email.UserName, "60%", "100%", "100%",
		email.InvoiceNumber, email.ClassName, "100%",
		email.ClassImage, "100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplatePaymentRevised(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(`
	<table width="%s">
	<tr>
		<td style="display:block!important;max-width:600px!important;clear:both!important;margin:0 auto;padding:0;background-color:transparent";>
			<div style="min-width:600px;display:block;border-collapse:collapse;margin:0 auto;border:1px solid #e7e7e7";>
			   <table style="border-spacing:0;width:%s;background-color:transparent;margin:0;padding:20px" cellspacing="0" cellpadding="0">
					<tbody>
						<tr style="margin:0;padding:0">
							<td style="text-align: center;">
								<img src="https://www.belajariah.com/email-assets/img/Icon-Belajariah.png" width="80px" alt="" style="margin:10px 0">
							</td>
							</tr>
							<tr>
								<td>
									<p style="font-size:14px;margin-bottom:10px">
										Assalamu’alaikum, <b style="color:#212121">%s</b>,
									</p>
								</td>
							</tr>
							<tr style="margin:0;padding:0">
								<td style="margin:0;padding:0">
									<h5 style="line-height:1.4;color:black;font-weight:700;margin:0px 0px 10px 0px;padding:0;font-size: 14px;">Mohon maaf pembelian kelas telah dibatalkan, silahkan kirimkan kembali bukti pembayarannya.</h5>  
								</td>
							</tr>
					</tbody>
				</table>
				<table style="padding:0 20px;margin-bottom:30px;margin: 0px auto 30px auto;width: %s;" cellspacing="0" cellpadding="0">
					<tbody>
						<tr>
							<td>
								<a href="#" style="background: linear-gradient(#542f91, #3f1c78, #835bc2);color: #fff;display:block;border: 2.5px solid #552b9c;border-radius:20px;padding:10px 10px;text-align:center;font-weight:600;text-decoration:none;font-size:14px">Kirim bukti pembayaran</a> 
							 </td>  
						</tr>
					</tbody>
				</table>
				<table cellspacing="0" cellpadding="0" style="width:%s">
					<tbody>
						<tr>
							<td style="padding:25px 20px 16px 20px">
								<h2 style="font-size:14px;font-weight:600;margin:0">Rincian Pesanan Kelas</h2>
							</td>
						</tr>
						<tr>
							<td style="padding-bottom:20px">
								<table style="border:0;padding:0 20px;width:%s" cellspacing="0" cellpadding="0">
									<tbody>
										<tr>
											<td colspan="2">
												<p style="margin:0 0 10px 0;">
														<span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">No. Invoice:</span>
														<span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s</span>
													</p>
												</td>
											</tr>
											<tr>
												<td colspan="2">
													<p style="margin:0 0 25px 0">
														<span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">Kelas:</span>
														<span style="margin-right:5px;color:rgba(49,53,59,0.96);font-size: 14px;">%s
													</p>
												</td>
											</tr>
											<tr>
												<td>
													<table cellspacing="0" cellpadding="0" style="width:%s">
														<tbody>
															<tr>
																<td style="text-align: center;">
																	<img src="%s" style="width:auto;height:180px;">
																</td>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr> 
					</tbody>
				</table>
				<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="margin:auto">
					<tbody>
						<tr>
							<td style="padding:40px 20px 0;background:#ffffff;font-family:sans-serif;font-size:12px;line-height:18px;color:#393d43">
								<p style="margin:0">E-mail ini dibuat otomatis, mohon tidak membalas. Jika butuh bantuan, silakan <a href="%s" style="text-decoration:none;color:#835bc2;font-weight:bold">hubungi CS Belajariah</a>.</p>
							</td>
						</tr>
						<tr>
							<td style="padding:32px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
									<tbody>
										<tr>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a"><p style="margin:0">Download Belajariah</p></td>
														</tr>
														<tr>
															<td style="padding:8px 0">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconGooglePlayStore.jpg" width="135" height="40" alt="banner" style="display:inline-block;width:"%s";height:40px;max-width:135px;margin:auto" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
											<td valign="top" width="%s">
												<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s">
													<tbody>
														<tr>
															<td style="font-family:sans-serif;font-size:12px;line-height:18px;color:#73767a;text-align:right">
																<p style="margin:0">Ikuti Kami</p>
															</td>
														</tr>
														<tr>
															<td style="padding:8px 0;text-align:right">
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconFb.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconYt.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
																<a href="%s" style="text-decoration:none">
																	<img src="https://www.belajariah.com/email-assets/img/IconIg.png" width="32" height="32" alt="banner" style="display:inline-block;height:32px;max-width:32px;margin:auto;" class="CToWUd">
																</a>
															</td>
														</tr>
													</tbody>
												</table>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
						<tr>
							<td style="padding:24px 20px 0;background:#ffffff">
								<table role="presentation" cellspacing="0" cellpadding="0" border="0" width="%s" style="border-top:1px solid #e5e7e9">
									<tbody>
										<tr>
											<td style="padding:4px 0 24px;font-family:sans-serif;font-size:12px;line-height:18px;color:#bdbec0;text-align:center">
												<p style="margin:0">%s</p>
											</td>
										</tr>
									</tbody>
								</table>
							</td>
						</tr>
					</tbody>
				</table> 
			</div>
		</td>
	</tr>
	</table>`,
		"100%", "100%", email.UserName, "60%", "100%", "100%",
		email.InvoiceNumber, email.ClassName, "100%",
		email.ClassImage, "100%",
		email.WhatsApp,
		"100%", "60%", "100%",
		email.GooglePLay,
		"100%", "40%", "100%",
		email.Facebook, email.Youtube, email.Instagram,
		"100%", email.CopyRight,
	)
	return bodyTemp
}

func TemplatePaymentCanceled(email model.EmailBody) string {
	bodyTemp := fmt.Sprintf(``)
	return bodyTemp
}
