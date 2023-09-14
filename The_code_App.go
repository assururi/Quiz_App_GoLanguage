
package main

import (
	"math/rand"
	"fmt"
	"time"
)


//untuk penyimpanan admin dan soal
type soal struct {
	question string
	choice [4]string
	answer int
	mudah int
}
type tabSoal struct {
	arrSoal [100]soal
	n int
}

//untuk menyimpan info nilai dan akun
type participant struct {
	userName string
	password string
	nilai float64
}

type tabParticipant struct {
	arrInfo [100]participant
	n int
}


func inputUser(T *tabParticipant){
	fmt.Print("USER NAME : ")
	fmt.Scan(&T.arrInfo[T.n].userName)
	fmt.Print("PASSWORD : ")
	fmt.Scan(&T.arrInfo[T.n].password)
	T.n = T.n + 1
}

func masukanSoal(K *tabSoal){
	K.arrSoal[0].question = "siapakah penemu gravitasi?"
	K.arrSoal[0].choice[0] = "Albert Einsten "
	K.arrSoal[0].choice[1] = "Isac Newton"
	K.arrSoal[0].choice[2] = "Nicola Tesla"
	K.arrSoal[0].choice[3] = "Thomas Edison"
	K.arrSoal[0].answer = 2
	K.arrSoal[1].question = "Negara manakah yang memiliki populasi terbanyak di dunia saat ini?"
	K.arrSoal[1].choice[0] = "India "
	K.arrSoal[1].choice[1] = "China"
	K.arrSoal[1].choice[2] = "Amerika Serikat"
	K.arrSoal[1].choice[3] = "Brazil"
	K.arrSoal[1].answer = 2
	K.arrSoal[2].question = "siapakah pelukis terkenal yang menciptakan Mona Lisa?"
	K.arrSoal[2].choice[0] = "Vincent Van Gogh "
	K.arrSoal[2].choice[1] = "Michelangelo"
	K.arrSoal[2].choice[2] = "Pablo Picasso"
	K.arrSoal[2].choice[3] = "Leanoardo da vinci"
	K.arrSoal[2].answer = 4
	K.arrSoal[3].question = "apa yang menjadi penyebab utama pemanasan global?"
	K.arrSoal[3].choice[0] = "emisi gas rumah kaca"
	K.arrSoal[3].choice[1] = "letusan gunung berapi"
	K.arrSoal[3].choice[2] = "aktivitas matahari"
	K.arrSoal[3].choice[3] = "perubahan orbit"
	K.arrSoal[3].answer = 1
	K.arrSoal[4].question = "apa nama ilmuan terkenal yang merupakan teori relativitas?"
	K.arrSoal[4].choice[0] = "Isaac Newton "
	K.arrSoal[4].choice[1] = "Albert Enisten"
	K.arrSoal[4].choice[2] = "Marie Curie"
	K.arrSoal[4].choice[3] = "Galileo galilie"
	K.arrSoal[4].answer = 2
	K.n = 5
}

func main(){
	var T tabParticipant
	var K tabSoal
	var user string
	masukanSoal(&K)
	for user != "DONE" {
		mainMenu(&T,&K, &user)
	}
	fmt.Println(T.arrInfo[0:T.n])
	fmt.Println(K.arrSoal[0:K.n])
}

func mainMenu(T *tabParticipant, K *tabSoal, user *string){
	fmt.Print(
		"---------------------------HOME MENU---------------------------\n",
		"information about menu : \n",
		"1    : untuk pergi ke admin menu \n",
		"2    : untuk pergi ke mahasiswa menu \n",
		"3    : untuk melihat nama dan nilai peserta \n",
		"DONE : untuk keluar \n",
		"----------------------------------------------------------------\n",
		"Your Answer : ",
	)
	fmt.Scan(user)
	if *user == "1"{
		fmt.Print("admin")
		menuAdmin(K)
	} else if *user == "2" {
		fmt.Print("menuParticipant")
		menuParticipant(T, K, *user)
	} else if *user == "3" {
		infopeserta(*T)
	}
}
func menuParticipant(T *tabParticipant, K *tabSoal, user string){
	// var user string
	// var K tabSoal
	fmt.Print(
		"pilih lah menu diabwah ini untuk melanjutkan ke page selanjutnya \n",
		"1 : untuk registerasi \n",
		"2 : untuk login\n",
		"3 : untuk back ke home menu \n",
		"---------------------------------------------------------------- \n",
		"Your Answer : ",
	)
	fmt.Scan(&user)
	if user == "1" {
		inputUser(T)
	} else if user == "2" {
		login(T, K)
	} else if user == "3" {
		mainMenu(T, K, &user)
	}
}

func found(T tabParticipant, name, pass string) bool {
	var hasil bool
	for i:=0;i<T.n;i++{
		if T.arrInfo[i].userName == name && T.arrInfo[i].password == pass {
			hasil = true
		}
	}
	return hasil
}

func login(T *tabParticipant, K *tabSoal){
	var name, pass, user string
	var ada bool
	fmt.Print(
		"---------------------------------------------------------------- \n",
		"------------------MASUKAN USERNAME DAN PASS ANDA---------------- \n",
		"---------------------------------------------------------------- \n",
	)
	fmt.Print("USERNAME : ")
	fmt.Scan(&name)
	fmt.Print("PASSWORD : ")
	fmt.Scan(&pass)
	ada = found(*T, name, pass)
	if ada {
		fmt.Println("Selamat anda telah masuk ke quiz")
		quizHomePage(T, K, name)

	} else {
		fmt.Print(
			"---------------------------------------------------------------- \n",
			"------maaf username dan password yang anda masukan salah-------- \n",
			"1 : untuk mencoba kembali \n",
			"2 : untuk kembali ke menu partisipasi\n",
			"----------------------------------------------------------------\n",
			"Your Answer : ",

		)
		fmt.Scan(&user)
		if user == "1"{
			login(T, K)
		} else if user == "2" {
			menuParticipant(T, K, user)
		}
	}
}
//tampilan quiz untuk peserta
func quizHomePage(T *tabParticipant, K *tabSoal, userName string){
	var user string
	fmt.Print(
		"---------------------------------------------------------------------------------- \n",
		"------------------SELAMAT ANDA TELAH MEMASUKI HALAMAN QUIZ :)--------------------- \n",
		"----------------------APAKAH ANDA INGIN MENEGRJAKAN QUIZ--------------------------\n",
		"Y = untuk melanjutkan \n",
		"N = untuk kembali ke menu halaman  \n",
		"---------------------------------------------------------------------------------- \n",
	)
	fmt.Print("Y atau N ? ")
	fmt.Scan(&user)
	if user == "Y" {
		fmt.Println("dapat langsung memulai quiz")
		halamanQuiz(T, K, userName)
	} else {
		fmt.Print()
	}
}
func shuffleQuestions(K *tabSoal){
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(K.n, func(i, j int){
		K.arrSoal[i], K.arrSoal[j] = K.arrSoal[j], K.arrSoal[i]
	})
}

// Searching : Binary Search
func mencariUsername(T tabParticipant, nama string, index *int) {
	var kr, kn int
	var fond bool
	kr = 0
	kn = *index - 1
	fond = false
	for kr <= kn && !fond {
		med := (kr + kn) / 2
		if T.arrInfo[med].userName < nama {
			kr = med + 1
		} else if T.arrInfo[med].userName > nama {
			kn = med
		} else {
			fond = true
			*index = med
		}
	}
}
//disini peserta sudah dapat mengerjakan soal quzi secara random
func halamanQuiz(T *tabParticipant, K *tabSoal, nama string){
	// fmt.Print(K.n)
	var idx int
	var jwbUser string
	mencariUsername(*T, nama, &idx)
	// fmt.Println(T.arrInfo[idx].userName)
	shuffleQuestions(K)
	T.arrInfo[idx].nilai = 0
	for i:=0;i<5;i++{
		for k:=0;k<2;k++{
			fmt.Println()
		}
		fmt.Println(i+1,".", K.arrSoal[i].question)
		fmt.Println("A. ", K.arrSoal[i].choice[0])
		fmt.Println("B. ", K.arrSoal[i].choice[1])
		fmt.Println("C. ", K.arrSoal[i].choice[2])
		fmt.Println("D. ", K.arrSoal[i].choice[3])
		fmt.Print("your answer : ")
		fmt.Scan(&jwbUser)
		if jwbUser == "A" {
			idxJ := 1
			if idxJ == K.arrSoal[i].answer {
				T.arrInfo[idx].nilai = T.arrInfo[idx].nilai + 20
				K.arrSoal[idx].mudah = K.arrSoal[idx].mudah + 1
			} 
		} else if jwbUser == "B" {
			idxJ := 2
			if idxJ == K.arrSoal[i].answer {
				T.arrInfo[idx].nilai = T.arrInfo[idx].nilai + 20
				K.arrSoal[idx].mudah = K.arrSoal[idx].mudah + 1
			}
		} else if jwbUser == "C" {
			idxJ := 3 
			if idxJ == K.arrSoal[i].answer {
				T.arrInfo[idx].nilai = T.arrInfo[idx].nilai + 20
				K.arrSoal[idx].mudah = K.arrSoal[idx].mudah + 1
			}
		} else if jwbUser == "D" {
			idxJ := 4
			if idxJ == K.arrSoal[i].answer {
				T.arrInfo[idx].nilai = T.arrInfo[idx].nilai + 20
				K.arrSoal[idx].mudah = K.arrSoal[idx].mudah + 1
			}
		}
	}
	tampilanSelesaiQuiz(*T, idx)
	fmt.Scan(&jwbUser)
	if jwbUser == "1" {	
		halamanQuiz(T, K, nama)
	} else {
		fmt.Print()
	}
}
func tampilanSelesaiQuiz(T tabParticipant, idx int){
	hasil := T.arrInfo[idx].nilai
	fmt.Print(
		"====================================================================\n",
		"------------------TERIMAKASIH TELAH MENGERJAKAN QUIZ---------------- \n",
		"-----------------------NILAI QUIZ ANDA ADALAH---------------------- \n",
		"---------------------------------",hasil,"-----------------------------\n",
		"===================================================================\n",
		"Apabila nilai anda kurang memuaskan, anda dapat mencobanya kemabali \n",
		"1 : untuk mencoba kembali\n",
		"2 : untuk kembali ke halaman utam\n",
		"Nilai anda adalah ", hasil, "\n",
		"--------------------------------------------------------------------\n",
		"Your Answer : ",
	)
}
//bagian menu admin
func menuAdmin(K *tabSoal){
	// var T tabParticipant
	var user string
	fmt.Print(
		"-----------------------WELCOME TO THE MENU ADMIN-------------------- \n",
		"information menu : \n",
		"1 : add a question to bankSoal \n",
		"2 : change or edit the question or answer in bankSoal \n",
		"3 : delete the question or answer in bankSoal \n",
		"4 : back home menu \n",
		"5 : info mengenai soal \n",
		"--------------------------------------------------------------------- \n",
		"------------------------input your answer below---------------------- \n",
		"Your Answer : ",
	)
	fmt.Scan(&user)
	if user == "1" {
		addSoal(K)
	} else if user == "2" {
		mengeditSoal(K)
	} else if user == "3" {
		deleteSoal(K)
	} else if user == "4" {
		fmt.Print()
	} else if user == "5" {
		soalExtrim(*K)
	}
}
func addSoal(K *tabSoal){
	fmt.Print(
		"selamat datang di halamn menambah soal, berikut panduannya : \n",
		"1. kalimat tidak dapat dipisahkan dengan sepasi, kami sarankan gunakan underscore '_'\n",
		"2. apabila telah selesai menambahkan soal, masukan kaliam 'STOP' pada masukan soal \n",
		"index kunci jawaban : \n",
		"pg1 untuk A\n",
		"pg2 untuk B\n",
		"pg3 untuk C\n",
		"pg4 untuk D\n",	
		"---------------------------------------------------------------------------------- \n",
	)
	i := K.n
	fmt.Print("Pertanyaan ke-", i+1, " ")
	fmt.Scan(&K.arrSoal[i].question)
	for K.arrSoal[i].question != "STOP"{
		for j:=0;j<4;j++{
			fmt.Print("pg-", j+1, " ")
			fmt.Scan(&K.arrSoal[i].choice[j])
		}
		fmt.Print("masukan kunci jawaban")
		fmt.Scan(&K.arrSoal[i].answer)
		K.n = K.n + 1
		i = i + 1
		fmt.Print("Pertnyaan ke-", i+1, " ")
		fmt.Scan(&K.arrSoal[i].question)
	}
}
func deleteSoal(K *tabSoal){
	var temp soal
	var idx int
	fmt.Print(
		"----------------------------------------------------------------------------------\n",
		"-----------------------------PANDUAN UNTUK MENGHAPUS SOAL------------------------- \n",
		"anda akan diminta untuk memasukan no soal yang akan dihapus \n",
		"soal akan ditampilkan dalam layar agar mempermudah anda untuk menghapus soal \n",
		"----------------------------------------------------------------------------------\n",
	)
	menampilkanSoal(*K)
	fmt.Print("masukan soal ke berapa yang ingin di hapus : ")
	fmt.Scan(&idx)
	idx = idx
	K.arrSoal[idx] = temp
	for idx < K.n {
		K.arrSoal[idx], K.arrSoal[idx + 1] = K.arrSoal[idx + 1], K.arrSoal[idx]
		idx = idx + 1
	} 
	K.n = K.n - 1
}
//mengedit soal
func mengeditSoal(K *tabSoal){
	var idx int
	var temp soal
	fmt.Print(
		"----------------------------------------------------------------------------------\n",
		"-----------------------------PANDUAN UNTUK MENGEDIT SOAL------------------------- \n",
		"anda akan diminta untuk memasukan no soal yang akan di-edit \n",
		"soal akan ditampilkan dalam layar agar mempermudah anda untuk mengedit soal dan jawabanya \n",
		"----------------------------------------------------------------------------------\n",
	)
	menampilkanSoal(*K)
	fmt.Print("Your Answer : ")
	fmt.Scan(&idx)
	fmt.Print(
		"-----------------------------Anda telah memilih soal------------------------------ \n",
		"------------------------TULIS PERTNYAAN BARU DI BAWAH INI------------------------- \n",
		"note : tidak menggunakan spasi (menggunakan garis bawah) \n",
		"Your Answer : ",
	)
	fmt.Scan(&temp.question)
	K.arrSoal[idx-1].question = temp.question
	fmt.Print(
		"-----------------------------Anda Berhasil Mengganti Soal------------------------------ \n",
		"--------------------------TULIS PILIHAN JAWABAN DI BAWAH INI--------------------------- \n",
		"note : tidak menggunakan spasi (menggunakan garis bawah) \n",
	)
	for i:=0;i<4;i++{
		fmt.Print(i+1, " ")
		fmt.Print(i+1, " ")
		fmt.Scan(&temp.choice[i])
		fmt.Println()
	}
	fmt.Print(
		"-----------------------Anda telah Mengubah Pilihan Ganda soal------------------------------ \n",
		"-----------------------TULIS KUNCI JAWABAN BARU DI BAWAH INI------------------------------- \n",
		"note : tidak menggunakan spasi (menggunakan garis bawah) \n",
		"       1 apabila A adalah benar\n",
		"		2 apabila B adalah benar \n",
		"		3 apabila C adalah benar \n",			
		"		4 apabila D adalah benar \n",
		"       langsung tuliskan jawabnya saja antara 1 sampai 4 \n",
		"------------------------------------------------------------------------------------------\n",
		"masuakn 1, 2, 3 atau 4 : ",
	)
	fmt.Print(&temp.answer)
}
func menampilkanSoal(K tabSoal){
	for i:=0;i<K.n;i++{
		fmt.Print(i + 1, ". ")
		fmt.Println(K.arrSoal[i].question)
		for k:=0;k<4;k++{
			fmt.Print("   PG", i + 1)
			fmt.Println(K.arrSoal[i].choice[k])
		}
	}
}
//menampilan soal yang termudah dan yang tersulit
func soalExtrim(K tabSoal){
	i, j := 0, 0
	var temp soal
	for i < K.n {
		j = i
		temp = K.arrSoal[j]
		for j > 0 && temp.mudah < K.arrSoal[j-1].mudah {
			K.arrSoal[j] = K.arrSoal[j-1]
			j = j - 1
		}
		K.arrSoal[j] = temp
		i = i + 1
	}
	fmt.Println("###soal mudah###")
	for z:=0;z<5;z++{
		fmt.Println(z+1, K.arrSoal[z].question)
	}
	//selection sort
	for l:=0;l<K.n;l++{
		min := l
		for m:=l+1;m<K.n;m++{
			if K.arrSoal[min].mudah > K.arrSoal[m].mudah {
				min = m
			}
		}
		K.arrSoal[l], K.arrSoal[min] = K.arrSoal[min], K.arrSoal[l]
	}
	fmt.Print("###soal sulit###")
	for z:=0;z<5;z++{
		fmt.Println(z+1, K.arrSoal[z].question)
	}
}

//mengelarkan nama dan semua siswa 
func infopeserta(T tabParticipant) {
	for i := 0; i < T.n; i++ {
		fmt.Println("Nama Peserta :", T.arrInfo[i].userName, ", Nilai Peserta :", T.arrInfo[i].nilai)
	}
}

