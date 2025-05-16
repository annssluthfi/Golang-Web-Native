package homecontroller

import "net/http"

// Parameter 'w' bertipe http.ResponseWriter, yaitu interface yang digunakan untuk mengirim respon ke client.
// Karena 'http.ResponseWriter' adalah interface, kita tidak perlu (dan tidak boleh) menggunakan tanda '*'.
// Go akan secara otomatis menangani objek sebenarnya di balik interface ini.

// Parameter 'r' bertipe *http.Request, yaitu pointer ke struct Request yang berisi semua data permintaan HTTP.
// Kita menggunakan pointer (*) agar lebih efisien (tidak menyalin seluruh struct), 
// dan agar kita bisa membaca atau memodifikasi data request secara langsung jika dibutuhkan.
func Welcome(w http.ResponseWriter, r *http.Request) {

}