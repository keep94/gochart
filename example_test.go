package gochart_test

import (
	"fmt"
	"math"
	"strings"

	"github.com/keep94/gochart"
	"github.com/keep94/gomath"
)

func ExampleFractionDigits() {
	xs := gochart.NewFloats(1.01, .01, 5)
	ys := xs.Apply(math.Sqrt)
	gochart.NewChart(xs, ys, gochart.FractionDigits(3, 4)).WriteTo(nil)
	// Output:
	// +-----+------+
	// |1.010|1.0050|
	// |1.020|1.0100|
	// |1.030|1.0149|
	// |1.040|1.0198|
	// |1.050|1.0247|
	// +-----+------+
}

func ExampleInts_ApplyBigInt() {
	// From the github.com/keep94/gomath package.
	p := gomath.NewPartition()
	xs := gochart.NewInts(1, 1, 100)
	ys := xs.ApplyBigInt(p.Chart)
	gochart.NewChart(xs, ys, gochart.NumRows(25)).WriteTo(nil)
	// Output:
	// +---+---------+---+---------+---+---------+---+---------+
	// |  1|        1| 26|     2436| 51|   239943| 76|  9289091|
	// |  2|        2| 27|     3010| 52|   281589| 77| 10619863|
	// |  3|        3| 28|     3718| 53|   329931| 78| 12132164|
	// |  4|        5| 29|     4565| 54|   386155| 79| 13848650|
	// |  5|        7| 30|     5604| 55|   451276| 80| 15796476|
	// |  6|       11| 31|     6842| 56|   526823| 81| 18004327|
	// |  7|       15| 32|     8349| 57|   614154| 82| 20506255|
	// |  8|       22| 33|    10143| 58|   715220| 83| 23338469|
	// |  9|       30| 34|    12310| 59|   831820| 84| 26543660|
	// | 10|       42| 35|    14883| 60|   966467| 85| 30167357|
	// | 11|       56| 36|    17977| 61|  1121505| 86| 34262962|
	// | 12|       77| 37|    21637| 62|  1300156| 87| 38887673|
	// | 13|      101| 38|    26015| 63|  1505499| 88| 44108109|
	// | 14|      135| 39|    31185| 64|  1741630| 89| 49995925|
	// | 15|      176| 40|    37338| 65|  2012558| 90| 56634173|
	// | 16|      231| 41|    44583| 66|  2323520| 91| 64112359|
	// | 17|      297| 42|    53174| 67|  2679689| 92| 72533807|
	// | 18|      385| 43|    63261| 68|  3087735| 93| 82010177|
	// | 19|      490| 44|    75175| 69|  3554345| 94| 92669720|
	// | 20|      627| 45|    89134| 70|  4087968| 95|104651419|
	// | 21|      792| 46|   105558| 71|  4697205| 96|118114304|
	// | 22|     1002| 47|   124754| 72|  5392783| 97|133230930|
	// | 23|     1255| 48|   147273| 73|  6185689| 98|150198136|
	// | 24|     1575| 49|   173525| 74|  7089500| 99|169229875|
	// | 25|     1958| 50|   204226| 75|  8118264|100|190569292|
	// +---+---------+---+---------+---+---------+---+---------+
}

func ExampleInts_ApplyBigIntStream() {
	// From the github.com/keep94/gomath package
	uglies := gomath.Ugly(3, 5, 7)
	xs := gochart.NewInts(1, 1, 100)
	ys := xs.ApplyBigIntStream(uglies)
	gochart.NewChart(xs, ys, gochart.NumCols(4)).WriteTo(nil)
	// Output:
	// +---+-----+---+-----+---+-----+---+-----+
	// |  1|    1| 26|  343| 51| 2625| 76|11025|
	// |  2|    3| 27|  375| 52| 2835| 77|11907|
	// |  3|    5| 28|  405| 53| 3087| 78|12005|
	// |  4|    7| 29|  441| 54| 3125| 79|13125|
	// |  5|    9| 30|  525| 55| 3375| 80|14175|
	// |  6|   15| 31|  567| 56| 3645| 81|15309|
	// |  7|   21| 32|  625| 57| 3675| 82|15435|
	// |  8|   25| 33|  675| 58| 3969| 83|15625|
	// |  9|   27| 34|  729| 59| 4375| 84|16807|
	// | 10|   35| 35|  735| 60| 4725| 85|16875|
	// | 11|   45| 36|  875| 61| 5103| 86|18225|
	// | 12|   49| 37|  945| 62| 5145| 87|18375|
	// | 13|   63| 38| 1029| 63| 5625| 88|19683|
	// | 14|   75| 39| 1125| 64| 6075| 89|19845|
	// | 15|   81| 40| 1215| 65| 6125| 90|21609|
	// | 16|  105| 41| 1225| 66| 6561| 91|21875|
	// | 17|  125| 42| 1323| 67| 6615| 92|23625|
	// | 18|  135| 43| 1575| 68| 7203| 93|25515|
	// | 19|  147| 44| 1701| 69| 7875| 94|25725|
	// | 20|  175| 45| 1715| 70| 8505| 95|27783|
	// | 21|  189| 46| 1875| 71| 8575| 96|28125|
	// | 22|  225| 47| 2025| 72| 9261| 97|30375|
	// | 23|  243| 48| 2187| 73| 9375| 98|30625|
	// | 24|  245| 49| 2205| 74|10125| 99|32805|
	// | 25|  315| 50| 2401| 75|10935|100|33075|
	// +---+-----+---+-----+---+-----+---+-----+
}

func ExampleInts_ApplySlice() {
	// From the github.com/keep94/gomath package
	factorsOf100 := gomath.Factors(100)
	xs := gochart.NewInts(1, 1, len(factorsOf100))
	ys := xs.ApplySlice(factorsOf100)
	gochart.NewChart(xs, ys).WriteTo(nil)
	// Output:
	// +-+---+
	// |1|  1|
	// |2|  2|
	// |3|  4|
	// |4|  5|
	// |5| 10|
	// |6| 20|
	// |7| 25|
	// |8| 50|
	// |9|100|
	// +-+---+
}

func ExampleInts_ApplyStream() {
	// From the github.com/keep94/gomath package
	harshads := gomath.Harshads(1)
	xs := gochart.NewInts(1, 1, 100)
	ys := xs.ApplyStream(harshads)
	gochart.NewChart(xs, ys, gochart.NumCols(4)).WriteTo(nil)
	// Output:
	// +---+---+---+---+---+---+---+---+
	// |  1|  1| 26| 63| 51|156| 76|252|
	// |  2|  2| 27| 70| 52|162| 77|261|
	// |  3|  3| 28| 72| 53|171| 78|264|
	// |  4|  4| 29| 80| 54|180| 79|266|
	// |  5|  5| 30| 81| 55|190| 80|270|
	// |  6|  6| 31| 84| 56|192| 81|280|
	// |  7|  7| 32| 90| 57|195| 82|285|
	// |  8|  8| 33|100| 58|198| 83|288|
	// |  9|  9| 34|102| 59|200| 84|300|
	// | 10| 10| 35|108| 60|201| 85|306|
	// | 11| 12| 36|110| 61|204| 86|308|
	// | 12| 18| 37|111| 62|207| 87|312|
	// | 13| 20| 38|112| 63|209| 88|315|
	// | 14| 21| 39|114| 64|210| 89|320|
	// | 15| 24| 40|117| 65|216| 90|322|
	// | 16| 27| 41|120| 66|220| 91|324|
	// | 17| 30| 42|126| 67|222| 92|330|
	// | 18| 36| 43|132| 68|224| 93|333|
	// | 19| 40| 44|133| 69|225| 94|336|
	// | 20| 42| 45|135| 70|228| 95|342|
	// | 21| 45| 46|140| 71|230| 96|351|
	// | 22| 48| 47|144| 72|234| 97|360|
	// | 23| 50| 48|150| 73|240| 98|364|
	// | 24| 54| 49|152| 74|243| 99|370|
	// | 25| 60| 50|153| 75|247|100|372|
	// +---+---+---+---+---+---+---+---+
}

func ExampleFloats_Apply() {
	xs := gochart.NewFloats(1.0, 1.0, 100)
	ys := xs.Apply(math.Sqrt)
	gochart.NewChart(
		xs, ys, gochart.NumCols(4), gochart.YFormat("%.4f")).WriteTo(nil)
	// Output:
	// +---+-------+---+-------+---+-------+---+-------+
	// |  1| 1.0000| 26| 5.0990| 51| 7.1414| 76| 8.7178|
	// |  2| 1.4142| 27| 5.1962| 52| 7.2111| 77| 8.7750|
	// |  3| 1.7321| 28| 5.2915| 53| 7.2801| 78| 8.8318|
	// |  4| 2.0000| 29| 5.3852| 54| 7.3485| 79| 8.8882|
	// |  5| 2.2361| 30| 5.4772| 55| 7.4162| 80| 8.9443|
	// |  6| 2.4495| 31| 5.5678| 56| 7.4833| 81| 9.0000|
	// |  7| 2.6458| 32| 5.6569| 57| 7.5498| 82| 9.0554|
	// |  8| 2.8284| 33| 5.7446| 58| 7.6158| 83| 9.1104|
	// |  9| 3.0000| 34| 5.8310| 59| 7.6811| 84| 9.1652|
	// | 10| 3.1623| 35| 5.9161| 60| 7.7460| 85| 9.2195|
	// | 11| 3.3166| 36| 6.0000| 61| 7.8102| 86| 9.2736|
	// | 12| 3.4641| 37| 6.0828| 62| 7.8740| 87| 9.3274|
	// | 13| 3.6056| 38| 6.1644| 63| 7.9373| 88| 9.3808|
	// | 14| 3.7417| 39| 6.2450| 64| 8.0000| 89| 9.4340|
	// | 15| 3.8730| 40| 6.3246| 65| 8.0623| 90| 9.4868|
	// | 16| 4.0000| 41| 6.4031| 66| 8.1240| 91| 9.5394|
	// | 17| 4.1231| 42| 6.4807| 67| 8.1854| 92| 9.5917|
	// | 18| 4.2426| 43| 6.5574| 68| 8.2462| 93| 9.6437|
	// | 19| 4.3589| 44| 6.6332| 69| 8.3066| 94| 9.6954|
	// | 20| 4.4721| 45| 6.7082| 70| 8.3666| 95| 9.7468|
	// | 21| 4.5826| 46| 6.7823| 71| 8.4261| 96| 9.7980|
	// | 22| 4.6904| 47| 6.8557| 72| 8.4853| 97| 9.8489|
	// | 23| 4.7958| 48| 6.9282| 73| 8.5440| 98| 9.8995|
	// | 24| 4.8990| 49| 7.0000| 74| 8.6023| 99| 9.9499|
	// | 25| 5.0000| 50| 7.0711| 75| 8.6603|100|10.0000|
	// +---+-------+---+-------+---+-------+---+-------+
}

func ExampleFloats_ApplyInv() {
	xs := gochart.NewFloats(1.0, 1.0, 300)
	ys := xs.ApplyInv(
		func(x float64) float64 {
			return math.Pow(x, x)
		},
		1.0,
		5.0)
	gochart.NewChart(
		xs, ys, gochart.NumRows(50), gochart.YFormat("%.4f")).WriteTo(nil)
	// Output:
	// +---+------+---+------+---+------+---+------+---+------+---+------+
	// |  1|1.0000| 51|3.2963|101|3.6016|151|3.7761|201|3.8981|251|3.9917|
	// |  2|1.5596| 52|3.3051|102|3.6060|152|3.7789|202|3.9002|252|3.9934|
	// |  3|1.8255| 53|3.3138|103|3.6102|153|3.7818|203|3.9023|253|3.9951|
	// |  4|2.0000| 54|3.3223|104|3.6145|154|3.7845|204|3.9044|254|3.9967|
	// |  5|2.1294| 55|3.3307|105|3.6187|155|3.7873|205|3.9064|255|3.9984|
	// |  6|2.2318| 56|3.3388|106|3.6228|156|3.7901|206|3.9085|256|4.0000|
	// |  7|2.3165| 57|3.3468|107|3.6269|157|3.7928|207|3.9105|257|4.0016|
	// |  8|2.3884| 58|3.3547|108|3.6310|158|3.7955|208|3.9126|258|4.0033|
	// |  9|2.4510| 59|3.3624|109|3.6350|159|3.7982|209|3.9146|259|4.0049|
	// | 10|2.5062| 60|3.3700|110|3.6390|160|3.8009|210|3.9166|260|4.0065|
	// | 11|2.5556| 61|3.3775|111|3.6429|161|3.8036|211|3.9186|261|4.0081|
	// | 12|2.6003| 62|3.3848|112|3.6468|162|3.8062|212|3.9206|262|4.0097|
	// | 13|2.6411| 63|3.3920|113|3.6507|163|3.8089|213|3.9226|263|4.0113|
	// | 14|2.6785| 64|3.3991|114|3.6546|164|3.8115|214|3.9246|264|4.0129|
	// | 15|2.7132| 65|3.4061|115|3.6584|165|3.8141|215|3.9266|265|4.0145|
	// | 16|2.7454| 66|3.4129|116|3.6621|166|3.8167|216|3.9285|266|4.0160|
	// | 17|2.7754| 67|3.4197|117|3.6659|167|3.8192|217|3.9305|267|4.0176|
	// | 18|2.8037| 68|3.4263|118|3.6696|168|3.8218|218|3.9324|268|4.0192|
	// | 19|2.8302| 69|3.4329|119|3.6732|169|3.8243|219|3.9344|269|4.0207|
	// | 20|2.8553| 70|3.4393|120|3.6769|170|3.8269|220|3.9363|270|4.0223|
	// | 21|2.8791| 71|3.4457|121|3.6805|171|3.8294|221|3.9382|271|4.0238|
	// | 22|2.9016| 72|3.4519|122|3.6840|172|3.8318|222|3.9401|272|4.0254|
	// | 23|2.9231| 73|3.4581|123|3.6876|173|3.8343|223|3.9420|273|4.0269|
	// | 24|2.9436| 74|3.4641|124|3.6911|174|3.8368|224|3.9439|274|4.0284|
	// | 25|2.9632| 75|3.4701|125|3.6946|175|3.8392|225|3.9458|275|4.0300|
	// | 26|2.9820| 76|3.4760|126|3.6980|176|3.8416|226|3.9476|276|4.0315|
	// | 27|3.0000| 77|3.4818|127|3.7015|177|3.8441|227|3.9495|277|4.0330|
	// | 28|3.0173| 78|3.4876|128|3.7049|178|3.8465|228|3.9513|278|4.0345|
	// | 29|3.0340| 79|3.4932|129|3.7082|179|3.8489|229|3.9532|279|4.0360|
	// | 30|3.0500| 80|3.4988|130|3.7116|180|3.8512|230|3.9550|280|4.0375|
	// | 31|3.0655| 81|3.5043|131|3.7149|181|3.8536|231|3.9568|281|4.0390|
	// | 32|3.0804| 82|3.5098|132|3.7182|182|3.8559|232|3.9587|282|4.0405|
	// | 33|3.0949| 83|3.5152|133|3.7214|183|3.8583|233|3.9605|283|4.0419|
	// | 34|3.1089| 84|3.5205|134|3.7247|184|3.8606|234|3.9623|284|4.0434|
	// | 35|3.1225| 85|3.5257|135|3.7279|185|3.8629|235|3.9641|285|4.0449|
	// | 36|3.1356| 86|3.5309|136|3.7311|186|3.8652|236|3.9659|286|4.0463|
	// | 37|3.1484| 87|3.5360|137|3.7342|187|3.8675|237|3.9676|287|4.0478|
	// | 38|3.1608| 88|3.5410|138|3.7374|188|3.8697|238|3.9694|288|4.0492|
	// | 39|3.1729| 89|3.5460|139|3.7405|189|3.8720|239|3.9712|289|4.0507|
	// | 40|3.1846| 90|3.5509|140|3.7436|190|3.8742|240|3.9729|290|4.0521|
	// | 41|3.1961| 91|3.5558|141|3.7466|191|3.8764|241|3.9747|291|4.0536|
	// | 42|3.2072| 92|3.5606|142|3.7497|192|3.8787|242|3.9764|292|4.0550|
	// | 43|3.2181| 93|3.5654|143|3.7527|193|3.8809|243|3.9781|293|4.0564|
	// | 44|3.2287| 94|3.5701|144|3.7557|194|3.8831|244|3.9799|294|4.0578|
	// | 45|3.2390| 95|3.5748|145|3.7587|195|3.8852|245|3.9816|295|4.0592|
	// | 46|3.2491| 96|3.5794|146|3.7616|196|3.8874|246|3.9833|296|4.0606|
	// | 47|3.2590| 97|3.5839|147|3.7646|197|3.8896|247|3.9850|297|4.0621|
	// | 48|3.2686| 98|3.5884|148|3.7675|198|3.8917|248|3.9867|298|4.0635|
	// | 49|3.2780| 99|3.5929|149|3.7704|199|3.8939|249|3.9884|299|4.0648|
	// | 50|3.2873|100|3.5973|150|3.7732|200|3.8960|250|3.9901|300|4.0662|
	// +---+------+---+------+---+------+---+------+---+------+---+------+
}

func ExampleChart_WriteTo() {
	xs := gochart.NewFloats(0.01, 0.01, 157)
	ys := xs.Apply(math.Sin)
	gochart.NewChart(
		xs,
		ys,
		gochart.NumRows(50),
		gochart.XFormat("%.2f"),
		gochart.YFormat("%.4f")).WriteTo(nil)
	// Output:
	// +----+------+----+------+----+------+----+------+
	// |0.01|0.0100|0.51|0.4882|1.01|0.8468|1.51|0.9982|
	// |0.02|0.0200|0.52|0.4969|1.02|0.8521|1.52|0.9987|
	// |0.03|0.0300|0.53|0.5055|1.03|0.8573|1.53|0.9992|
	// |0.04|0.0400|0.54|0.5141|1.04|0.8624|1.54|0.9995|
	// |0.05|0.0500|0.55|0.5227|1.05|0.8674|1.55|0.9998|
	// |0.06|0.0600|0.56|0.5312|1.06|0.8724|1.56|0.9999|
	// |0.07|0.0699|0.57|0.5396|1.07|0.8772|1.57|1.0000|
	// |0.08|0.0799|0.58|0.5480|1.08|0.8820|    |      |
	// |0.09|0.0899|0.59|0.5564|1.09|0.8866|    |      |
	// |0.10|0.0998|0.60|0.5646|1.10|0.8912|    |      |
	// |0.11|0.1098|0.61|0.5729|1.11|0.8957|    |      |
	// |0.12|0.1197|0.62|0.5810|1.12|0.9001|    |      |
	// |0.13|0.1296|0.63|0.5891|1.13|0.9044|    |      |
	// |0.14|0.1395|0.64|0.5972|1.14|0.9086|    |      |
	// |0.15|0.1494|0.65|0.6052|1.15|0.9128|    |      |
	// |0.16|0.1593|0.66|0.6131|1.16|0.9168|    |      |
	// |0.17|0.1692|0.67|0.6210|1.17|0.9208|    |      |
	// |0.18|0.1790|0.68|0.6288|1.18|0.9246|    |      |
	// |0.19|0.1889|0.69|0.6365|1.19|0.9284|    |      |
	// |0.20|0.1987|0.70|0.6442|1.20|0.9320|    |      |
	// |0.21|0.2085|0.71|0.6518|1.21|0.9356|    |      |
	// |0.22|0.2182|0.72|0.6594|1.22|0.9391|    |      |
	// |0.23|0.2280|0.73|0.6669|1.23|0.9425|    |      |
	// |0.24|0.2377|0.74|0.6743|1.24|0.9458|    |      |
	// |0.25|0.2474|0.75|0.6816|1.25|0.9490|    |      |
	// |0.26|0.2571|0.76|0.6889|1.26|0.9521|    |      |
	// |0.27|0.2667|0.77|0.6961|1.27|0.9551|    |      |
	// |0.28|0.2764|0.78|0.7033|1.28|0.9580|    |      |
	// |0.29|0.2860|0.79|0.7104|1.29|0.9608|    |      |
	// |0.30|0.2955|0.80|0.7174|1.30|0.9636|    |      |
	// |0.31|0.3051|0.81|0.7243|1.31|0.9662|    |      |
	// |0.32|0.3146|0.82|0.7311|1.32|0.9687|    |      |
	// |0.33|0.3240|0.83|0.7379|1.33|0.9711|    |      |
	// |0.34|0.3335|0.84|0.7446|1.34|0.9735|    |      |
	// |0.35|0.3429|0.85|0.7513|1.35|0.9757|    |      |
	// |0.36|0.3523|0.86|0.7578|1.36|0.9779|    |      |
	// |0.37|0.3616|0.87|0.7643|1.37|0.9799|    |      |
	// |0.38|0.3709|0.88|0.7707|1.38|0.9819|    |      |
	// |0.39|0.3802|0.89|0.7771|1.39|0.9837|    |      |
	// |0.40|0.3894|0.90|0.7833|1.40|0.9854|    |      |
	// |0.41|0.3986|0.91|0.7895|1.41|0.9871|    |      |
	// |0.42|0.4078|0.92|0.7956|1.42|0.9887|    |      |
	// |0.43|0.4169|0.93|0.8016|1.43|0.9901|    |      |
	// |0.44|0.4259|0.94|0.8076|1.44|0.9915|    |      |
	// |0.45|0.4350|0.95|0.8134|1.45|0.9927|    |      |
	// |0.46|0.4439|0.96|0.8192|1.46|0.9939|    |      |
	// |0.47|0.4529|0.97|0.8249|1.47|0.9949|    |      |
	// |0.48|0.4618|0.98|0.8305|1.48|0.9959|    |      |
	// |0.49|0.4706|0.99|0.8360|1.49|0.9967|    |      |
	// |0.50|0.4794|1.00|0.8415|1.50|0.9975|    |      |
	// +----+------+----+------+----+------+----+------+
}

func ExampleSquare() {
	xs := gochart.NewInts(1, 1, 10)
	ys := xs.Apply(func(x int64) int64 { return x * x })
	var builder strings.Builder
	n, err := gochart.NewChart(xs, ys).WriteTo(&builder)
	fmt.Println(n, err)
	fmt.Println(builder.String())
	// Output:
	// 108 <nil>
	// +--+---+
	// | 1|  1|
	// | 2|  4|
	// | 3|  9|
	// | 4| 16|
	// | 5| 25|
	// | 6| 36|
	// | 7| 49|
	// | 8| 64|
	// | 9| 81|
	// |10|100|
	// +--+---+
}
