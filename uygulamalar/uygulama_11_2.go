// uygulama_11_2
package main

import (
	"fmt"
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func main() {
	// veri setini yükle
	rawData, err := base.ParseCSVToInstances("iris.csv", false)
	if err != nil {
		panic(err)
	}
	// fmt.Println(rawData) // Ham veri özetlerini yazdır
	// Öklid mesafesi ile KNN sınıflandırıcısı kullan
	k := 3 // kaç komşuya bakılsın
	küme := knn.NewKnnClassifier("euclidean", k)
	//Veri setini iki parçaya (0.5) böl:
	//egitimVerisi ve testVerisi
	egitimVerisi, testVerisi := base.InstancesTrainTestSplit(rawData, 0.5)
	küme.Fit(egitimVerisi)
	//Test verisindeki her nokta için Öklid uzaklığını
	// hesapla ve en uygun sınıfı bul. tahminler değişkenine
	// sınıflandırma sonuçlarını aktar
	tahminler, _ := küme.Predict(testVerisi)
	fmt.Println("\nKNN bitti, sonuçlara bakalım:\n")
	// Karşılaştırma sonuçlarını ekrana yaz
	Karsilastirma, _ := evaluation.GetConfusionMatrix(testVerisi, tahminler)
	fmt.Println(evaluation.GetSummary(Karsilastirma))
	fmt.Println("Tahminlerin doğruluk oranı: ", evaluation.GetAccuracy(Karsilastirma))
}
