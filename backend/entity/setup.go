package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Doctor{}, &PatientType{}, &Symptoms{}, &Patient{},
	)

	db = database

	Phonsak := Doctor{
		DoctorName:       "Phonsak Songsang",
		DoctorEmail:      "Phon@gmail.com",
		DoctorPassword:   "Phonsak01",
		DoctorDepartment: "H. Surgery",
	}
	db.Model(&Doctor{}).Create(&Phonsak)

	Hanoi := Doctor{
		DoctorName:       "Hanoi Slotmachine",
		DoctorEmail:      "Hanoi@hotmail.in.th",
		DoctorPassword:   "Hanoiploy",
		DoctorDepartment: "h. Surgery",
	}
	db.Model(&Doctor{}).Create(&Hanoi)

	Kanok := Doctor{
		DoctorName:       "Kanokthip Lamai",
		DoctorEmail:      "Kanok@hotmail.com",
		DoctorPassword:   "Pookpik05",
		DoctorDepartment: "Surgery",
	}
	db.Model(&Doctor{}).Create(&Kanok)
	// ตาราง type.....................................................
	type1 := PatientType{
		Type: "outpatient",
	}
	db.Model(&PatientType{}).Create(&type1)

	type2 := PatientType{
		Type: "inpatient",
	}
	db.Model(&PatientType{}).Create(&type2)
	//ตารางโรค.......................................................
	Sym1 := Symptoms{
		SymptomsName: "CRF",
	}
	db.Model(&Symptoms{}).Create(&Sym1)

	Sym2 := Symptoms{
		SymptomsName: "Hypertension",
	}
	db.Model(&Symptoms{}).Create(&Sym2)

	Sym3 := Symptoms{
		SymptomsName: "ICH",
	}
	db.Model(&Symptoms{}).Create(&Sym3)

	Sym4 := Symptoms{
		SymptomsName: "AF",
	}
	db.Model(&Symptoms{}).Create(&Sym4)

	//ตารางผู้ป่วย.......................................................
	db.Model(&Patient{}).Create(&Patient{
		PatientsName: "Arnon Derek",
		DateAdmit:    time.Date(2022, 1, 2, 9, 0, 0, 0, time.Now().Location()),
		Age:          32,
		Doctor:       Phonsak,
		Symptoms:     Sym3,
		PatientType:  type2,
	})

	db.Model(&Patient{}).Create(&Patient{
		PatientsName: "Darin Darwin",
		DateAdmit:    time.Date(2022, 3, 2, 10, 0, 0, 0, time.Now().Location()),
		Age:          51,
		Doctor:       Hanoi,
		Symptoms:     Sym1,
		PatientType:  type1,
	})

	db.Model(&Patient{}).Create(&Patient{
		PatientsName: "Thana Ngampon",
		DateAdmit:    time.Date(2022, 2, 3, 9, 0, 0, 0, time.Now().Location()),
		Age:          60,
		Doctor:       Phonsak,
		Symptoms:     Sym4,
		PatientType:  type2,
	})

	// Q
	// var target Doctor
	// db.Model(&Doctor{}).Find(&target, db.Where("email = ?", "Phon@gmail.com"))

	// var watchedPlaylist Playlist
	// db.Model(&Playlist{}).Find(&watchedPlaylist, db.Where("title = ? and owner_id = ?", "Watched", target.ID))

	// var watchedList []*WatchVideo
	// db.Model(&WatchVideo{}).
	// 	Joins("Playlist").
	// 	Joins("Resolution").
	// 	Joins("Video").
	// 	Find(&watchedList, db.Where("playlist_id = ?", watchedPlaylist.ID))

	// for _, wl := range watchedList {
	// 	fmt.Printf("Watch Video: %v\n", wl.ID)
	// 	fmt.Printf("%v\n", wl.Playlist.Title)
	// 	fmt.Printf("%v\n", wl.Resolution.Value)
	// 	fmt.Printf("%v\n", wl.Video.Name)
	// 	fmt.Println("====")
	// }
}
