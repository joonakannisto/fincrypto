package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"sort"
)

func main() {
	finSyllab := map[string]int{"ta": 1064, "si": 624, "a": 599, "sa": 539, "ti": 536, "tä": 434, "li": 405, "o": 337, "te": 328, "tu": 298, "nen": 291, "la": 270, "e": 266, "to": 252, "ko": 245, "ki": 244, "le": 243, "mi": 242, "ke": 241, "ri": 219, "sen": 215, "ja": 210, "sä": 190, "lä": 181, "va": 177, "ka": 170, "taa": 156, "den": 151, "vat": 151, "na": 146, "ni": 145, "se": 145, "ä": 143, "lu": 142, "lun": 142, "lai": 136, "ra": 131, "ty": 131, "pi": 129, "ses": 129, "ku": 127, "nut": 125, "pe": 124, "kin": 123, "pa": 123, "vi": 117, "ma": 116, "toi": 113, "me": 111, "val": 111, "ten": 109, "taan": 107, "vä": 107, "nä": 106, "maan": 105, "an": 103, "kan": 100, "so": 99, "set": 98, "y": 93, "kun": 91, "tiin": 91, "ai": 87, "kä": 86, "pu": 83, "nis": 82, "lo": 81, "nyt": 81, "ve": 81, "tus": 80, "vät": 80, "sel": 79, "no": 77, "van": 77, "tet": 76, "vas": 76, "kai": 75, "tää": 75, "yh": 75, "kes": 73, "sin": 73, "vaa": 73, "pää": 72, "en": 70, "ne": 70, "ha": 68, "ky": 68, "mis": 68, "lis": 67, "i": 66, "jen": 66, "lin": 66, "man": 66, "rin": 65, "suu": 65, "tar": 65, "tään": 64, "tai": 63, "tei": 63, "nan": 62, "jat": 61, "kas": 60, "lan": 60, "saa": 60, "tun": 60, "maa": 59, "del": 58, "on": 58, "seen": 58, "kau": 57, "ken": 57, "po": 57, "tuk": 57, "jo": 56, "puo": 56, "tie": 56, "vai": 56, "kuu": 55, "mo": 55, "tel": 55, "kir": 54, "lii": 54, "mat": 54, "al": 53, "jan": 53, "teen": 53, "joh": 52, "mai": 52, "ro": 50, "tuu": 50, "hal": 49, "min": 49, "neet": 49, "suo": 49, "des": 48, "ar": 47, "il": 47, "jä": 47}
	finFirstSyllab := map[string]int{"o": 259, "a": 207, "ei": 196, "pa": 119, "pe": 111, "va": 109, "sa": 104, "ta": 101, "i": 93, "tu": 89, "toi": 80, "val": 80, "si": 78, "to": 77, "ko": 74, "kan": 71, "so": 71, "mie": 69, "y": 66, "ku": 61, "ai": 60, "us": 60, "pu": 59, "te": 59, "tie": 59, "ra": 56, "lu": 55, "tuo": 55, "van": 55, "pal": 54, "vä": 54, "py": 54, "voi": 53, "vii": 52, "ju": 51, "tar": 51, "en": 50, "pi": 50, "saa": 49, "vi": 48, "kau": 48, "kuu": 48, "tun": 48, "ha": 45, "yh": 44, "vas": 44, "ka": 43, "kä": 43, "muu": 43, "mi": 42, "pää": 42, "kir": 41, "seu": 41, "vuo": 40, "ve": 40, "syn": 40, "an": 39, "ar": 39, "oi": 39, "u": 38, "huo": 38, "ih": 37, "il": 37, "jo": 37, "kes": 37, "me": 37, "ym": 37, "on": 36, "hä": 35, "hy": 34, "ke": 34, "luo": 34, "po": 34, "vie": 33, "it": 33, "ni": 33, "nä": 33, "puo": 33, "se": 33, "hen": 32, "sel": 32, "suo": 32, "ker": 31, "läh": 31, "suu": 31, "kas": 30, "sy": 30, "joh": 29, "jou": 29, "kor": 29, "pai": 29, "uh": 29, "he": 28, "jäl": 28, "kat": 28, "kul": 28, "lä": 28, "maa": 28, "tai": 27, "hal": 26, "koh": 26, "mö": 26, "mu": 26, "mää": 26, "nou": 26, "ru": 26, "sil": 25, "su": 25, "ti": 25, "kun": 24, "vel": 24, "kuo": 23, "käy": 23, "vaa": 23, "tur": 22, "hei": 22, "käs": 22, "lau": 22, "ää": 22, "pel": 22, "vir": 22, "tul": 22, "kah": 21, "las": 21, "liit": 21, "0t": 21, "päi": 21, "rik": 20, "ri": 20, "aa": 19, "kään": 19, "lei": 19, "löy": 19, "mah": 19, "polt": 19, "tah": 19, "le": 18, "uu": 18}
	firstTable := cumMap(finFirstSyllab)
	syllabTable := cumMap(finSyllab)
	sylmax := uint64(18021)
	firstMax := uint64(6164)
	entropy_counter := uint64(1)
	added_entropy := uint64(1)
	// We pass 64 bits of entropy and are finished
	for loop := 0; entropy_counter <= entropy_counter*added_entropy; loop++ {
		entropy_counter *= added_entropy
		randInt := uint64(65535)
		UintByte := make([]byte, 8)
		for randInt > 3*sylmax && loop > 0 || randInt > 10*firstMax {
			rByte := make([]byte, 2)
			_, err := rand.Read(rByte)
			if err != nil {
				panic(err)
			}
			UintByte[0] = rByte[0]
			UintByte[1] = rByte[1]
			randInt = binary.LittleEndian.Uint64(UintByte)
		}
		if loop > 0 {
			randInt = randInt % sylmax
			keys := make([]int, len(syllabTable))
			i := 0
			for k, _ := range syllabTable {
				keys[i] = k
				i++
			}
			sort.Ints(keys)
			for _, key := range keys {
				if randInt < uint64(key) {
					added_entropy =
						sylmax / uint64(finSyllab[syllabTable[key]])
					fmt.Printf("%s", syllabTable[key])
					break
				}
			}
		} else {
			randInt = randInt % firstMax
			keys := make([]int, len(firstTable))
			i := 0
			for k, _ := range firstTable {
				keys[i] = k
				i++
			}
			sort.Ints(keys)
			for _, key := range keys {
				if randInt < uint64(key) {
					added_entropy =
						firstMax / uint64(finFirstSyllab[firstTable[key]])
					fmt.Printf("%s", firstTable[key])
					break
				}
			}
		}

	}
	fmt.Printf("\n")

}

func cumMap(mappi map[string]int) map[int]string {
	palautus := make(map[int]string, len(mappi))
	vali := make(map[int]string, len(mappi))
	var keys []int
	var summat []int
	summat = append(summat, 0)
	for key, value := range mappi {
		vali[value] = key
		summat = append(summat, value)
		keys = append(keys, value)
	}

	sort.Ints(keys)
	for i, k := range keys {
		summat[i+1] = summat[i] + summat[i+1]
		palautus[summat[i+1]] = vali[k]
	}
	return palautus
}
