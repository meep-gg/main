package api

import . "meep.gg/protos/riot/account/v1"

var RiotAPIAddr string = ":50051"

var TestAccount = &RiotAccount{
	Puuid:    "9WX5pU7pGi8EQvSFfC5hsiDJRpqdvmp1Lf57qbq1zSpSIiz38oakVBcFR_dMC2Rd6nbYf4n_pLI8sw",
	GameName: "Leong",
	TagLine:  "NA1",
}

var PuuidList = []string{
	"9WX5pU7pGi8EQvSFfC5hsiDJRpqdvmp1Lf57qbq1zSpSIiz38oakVBcFR_dMC2Rd6nbYf4n_pLI8sw",
	"6mFcYwk80OfQbusmgwl_-OzE8PMvIZkM1hY8_ZYkd0S8--EFl1ZV1_iICSQZcVbM8qHFt8JD3Oek1g",
	"AHCZg-98xSDZP1YFdOr5pVRfmgRpjQstnYo3zMIlhwNtUWeh3WXQS3RunCfG4y1iQzNAwmbLm8ogHA",
	"SYDw0mLpmwma9WajGM7XoX4wyoCg_2RiweavZzz42V8XS8HRG3PDQB8Z-SyCj0lCpkZQ8JXCz8jNrQ",
	"qCSPKeQoHWqk2Q0sSHxmfZ2cMPZApUgn4ssQpuCgb1CwKq0TGxSCuK8pHLaBt5PhbUxISMMDWiu4VA",
	"rnjVDsf4fe1jE4qMsh3SUo-ZfzXJZ953utjUGRhPoJJP6HpAEdJ2uVpSL7MSN_GAurKSlqcC5uPHLg",
	"MizP75iCKP6n3gGAHmkKUU9aMiqcb4KztnXeY51xhoJ_rKm7x-9DKNV0iGqCi9TQxgGfd0Smu0aPWA",
	"d5VvtIVpEUjAcURd-EUg7_8npzgm2Xq_MDLolHHUxki30VKsicLS-2OKUlETQaMQ0zUel52d-fr83Q",
	"z0ovcYsnBVcTU1Pno0pABUkR53uiIop2mNnfrAbQXi4_lmtjRJM9KNxoFzgv8G808I4_zaIASAbCnQ",
	"6QIKeLgpuKyb3437FbjKOG0_y3n4MeyT7tm3SQKN9iLNxCI1qPVSogXfuntVkil_5ZEi532IZ0keAA",
	"uJuu5k-8EOApqkmk-V069gFtYJ3ZVm_jiRWtlXrZFOCXoBoxJEjojkbkzZKGTAsc58y7ylm3116JaQ",
	"iYR4bLVm5DSThmYZdb67QArfGFsuXUu2UdVH4vAnRFk46vqF4YLMqXUyjrsmxNiXajvaCDkqhsga9w",
	"_hY6EbdV0VInMNl89igP8GZ6EZXJKJUMixF9VNaOOn0YNPVYoSFoLn7R4i06AcwEbFrNtJAWSGuPog",
	// "nLlAzvI8dH3M2GQXpGuL2Q4kO3NwcDn1UqWyIvN_lNSqnwezlZui2GXJfCQ4D-5vYPIInL0uC9pNwA",
	// "t7s2-qVOuTPoaSJcwuX3MK3j581esr7DkIYRaGHwMk6Br3dG4KbXkPs05g7QQpSqvp3pLpj_JfoWPw",
	// "J5XZd28WIFV1nnaX22Ssr5NSG5-AeBbLe-188LtDeZXpo2jEVc9CaeFR9dpt8f96QfinJdfTKEkfYw",
	// "tBhqwv0fC4sMPd-YenN7IYVhvKwkRkGumvprz24to-5X69XmIlg4Q8b0NpZvrq4D9w1SZXYbbwTOmQ",
	// "EN3HXRifVwAIjrIAc-UjiZumqFX79pGWFAdcWoWK1uRdU2pphuLz-VqK7MRnT67ZErwhCgzCx2x3Zw",
	// "aQUU-4un297beXu8JCo8TO5aMdvJOdyNdWYW31zCtw-I4RqFLIHc91jOdSRAVVM1uGZD0-Q6lk8VrQ",
	// "oNFPeCEgg0-gqdP61c51GoW2sJsqOEAc5L32t-0f6Vi6qr1Cl7T_-ajvaTE4kZW-eaqoT3jKgXPwDA",
	// "KKxoerRGOsHiuxM5XtlNfnjUQFulmxK3OvAoTqk6WmwM56-gM6R0fVkU4pldO6yL0hlgY9n5B_aZfA",
	// "OYYXqhJzvdY7OTFfp5jCqpmhrX_LV8BZFXbdenXBFSk0634ZDB1ANZ1DGPU_xfuf1kBUMVT2KinKPw",
	// "QA0Udpu_7mauqE14YpsoH1ycnTuFm374CmqjgPqN_PRr_IZyfPVOYcChG3YuyXuTGGEEzJEncR4MJw",
	// "ZD0SAhDcK0veSeL2fepFu0We7UoviuDHBHGjmHa7me7-t99V3pMhn8oIW8g7JtJ70sbkEH-AWXBBmg",
	// "V6fEdgpQic6V8RG8lFeTHaYXLklBMe6RhizTohqWh54rQT7M-Qo5aKtCNaTgzzJcj03MymaapbIaIg",
	// "UQVKERbHeucJUDIC9sdsCS84MKNKabrYnkw59GV2_nEFQXZvGzixCC4zrMZ4AzTqNwS8pukdSPZ2uA",
	// "qcv2kQC-eEbgJSEYY_ksstW8OdHKOrVf_qAiRw7r--r9Qs6lUd6n_Av8NfOSonhr2A-mUi8ZKdEzLA",
	// "D074vTrUfKFyndgwyS2nQcJO_ZNdfmg8fTPJSYPwpFU3H_jMoT97wXXuRsbiSLMl8QozHTuatD2fRw",
}
