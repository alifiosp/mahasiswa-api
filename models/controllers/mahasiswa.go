package controllers

import (
	"net/http"

	"github.com/alifiosp/mahasiswa-api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type MahasiswaInput struct {
	Nim  string `json:"nim"`
	Nama string `json:"nama"`
}

// view data
func Tampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(200, gin.H{
		"data": mhs,
	})
}

//add Data[POST]
func MahasiswaTambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi inputan/masukan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Data Salah",
		})
		return
	}

	//proses input data
	mhs := models.Mahasiswa{
		Nim:  dataInput.Nim,
		Nama: dataInput.Nama,
	}

	//membuat data
	db.Create(&mhs)

	//menampilkan hasil
	c.JSON(200, gin.H{
		"data": mhs,
	})
}

//ubah data berdasarkan nim
func MahasiswaUbah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// cek dulu data ada atau tidak
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data mahasiswa tidak ditemukan",
		})
		return
	}

	//validasi inputan/masukan
	var dataInput MahasiswaInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Data Salah",
		})
		return
	}

	//ubah data
	db.Model(&mhs).Update(dataInput)

	//menampilkan hasil
	c.JSON(200, gin.H{
		"data": mhs,
	})
}

//hapus data
func MahasiswaHapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	// cek dulu data ada atau tidak
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "data mahasiswa tidak ditemukan",
		})
		return
	}

	//ubah data
	db.Delete(&mhs)

	//menampilkan hasil
	c.JSON(200, gin.H{
		"data": "berhasil dihapus",
	})
}
