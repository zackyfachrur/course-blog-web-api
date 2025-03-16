package models

// Album struct
type Album struct {
	ID        string `json:"id"`
	Judul     string `json:"judul"`
	Deskripsi string `json:"deskripsi"`
	Thumbnail string `json:"thumbnail"`
	Kategori  string `json:"kategori"`
	Harga     string `json:"harga"`
}

// Data dummy album
var Albums = []Album{
	{
		ID:        "1",
		Judul:     "Dasar-Dasar Manajemen Waktu yang Efektif",
		Deskripsi: "Pelajari teknik mengatur waktu agar lebih produktif dan efisien.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-1.jpg",
		Kategori:  "pengembanganDiri",
		Harga:     "150000",
	},
	{
		ID:        "2",
		Judul:     "Strategi Digital Marketing",
		Deskripsi: "Pelajari dasar-dasar digital marketing untuk meningkatkan penjualan online.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-2.jpg",
		Kategori:  "pemasaran",
		Harga:     "250000",
	},
	{
		ID:        "3",
		Judul:     "Membangun Mindset Sukses dalam",
		Deskripsi: "Temukan cara membangun pola pikir positif untuk mencapai kesuksesan.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-3.jpg",
		Kategori:  "pengembanganDiri",
		Harga:     "180000",
	},
	{
		ID:        "4",
		Judul:     "Rahasia Sukses Menjalankan Bisnis Online",
		Deskripsi: "Pelajari strategi penting untuk memulai dan mengembangkan bisnis online.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-4.jpg",
		Kategori:  "bisnis",
		Harga:     "275000",
	},
	{
		ID:        "5",
		Judul:     "Personal Branding di Era Digital",
		Deskripsi: "Bangun citra diri yang kuat dan profesional di media sosial.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-5.jpg",
		Kategori:  "pengembanganDiri",
		Harga:     "220000",
	},
	{
		ID:        "6",
		Judul:     "Copywriting yang Menjual untuk Bisnis Online",
		Deskripsi: "Teknik menulis promosi yang efektif untuk menarik pelanggan.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-6.jpg",
		Kategori:  "pemasaran",
		Harga:     "230000",
	},
	{
		ID:        "7",
		Judul:     "Membuat Desain UI/UX Aplikasi Mobile dengan Figma",
		Deskripsi: "Panduan praktis dalam mendesain antarmuka aplikasi dengan Figma.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-7.jpg",
		Kategori:  "desain",
		Harga:     "280000",
	},
	{
		ID:        "8",
		Judul:     "Manajemen Keuangan untuk Pebisnis Pemula",
		Deskripsi: "Belajar cara mengatur keuangan bisnis agar tetap stabil.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-8.jpg",
		Kategori:  "bisnis",
		Harga:     "260000",
	},
	{
		ID:        "9",
		Judul:     "Strategi Pemasaran Konten yang Efektif",
		Deskripsi: "Pelajari cara membuat konten yang menarik dan relevan untuk audiens Anda.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-9.jpg",
		Kategori:  "pemasaran",
		Harga:     "200000",
	},
	{
		ID:        "10",
		Judul:     "Facebook Ads untuk Pemula",
		Deskripsi: "Pelajari cara membuat iklan Facebook yang efektif dan efisien.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-1.jpg",
		Kategori:  "pemasaran",
		Harga:     "220000",
	},
	{
		ID:        "11",
		Judul:     "Membangun Brand yang Kuat",
		Deskripsi: "Panduan dalam membangun identitas brand yang melekat di benak pelanggan.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-2.jpg",
		Kategori:  "pemasaran",
		Harga:     "300000",
	},
	{
		ID:        "12",
		Judul:     "Email Marketing yang Mengonversi",
		Deskripsi: "Cara menulis email yang menarik perhatian dan mendorong tindakan.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-3.jpg",
		Kategori:  "pemasaran",
		Harga:     "180000",
	},
	{
		ID:        "13",
		Judul:     "Pemasaran di TikTok untuk Bisnis",
		Deskripsi: "Strategi pemasaran di TikTok untuk meningkatkan brand awareness.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-4.jpg",
		Kategori:  "pemasaran",
		Harga:     "240000",
	},
	{
		ID:        "14",
		Judul:     "Instagram Marketing untuk UMKM",
		Deskripsi: "Cara memanfaatkan Instagram untuk meningkatkan penjualan bisnis kecil.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-5.jpg",
		Kategori:  "pemasaran",
		Harga:     "210000",
	},
	{
		ID:        "15",
		Judul:     "Growth Hacking untuk Startup",
		Deskripsi: "Teknik pemasaran cepat untuk pertumbuhan bisnis startup.",
		Thumbnail: "https://raw.githubusercontent.com/zackyfachrur/course-blog-web-app/refs/heads/main/public/thumbnail/video-thumb-6.jpg",
		Kategori:  "pemasaran",
		Harga:     "320000",
	},
}
