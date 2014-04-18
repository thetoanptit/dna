package am

import (
	"dna"
	"time"
)

func ExampleAPIAlbum() {
	album, err := GetAPIAlbum(243930)
	if err != nil {
		dna.Log(err.Error())
	} else {
		album.Checktime = time.Date(2013, time.November, 21, 0, 0, 0, 0, time.UTC)
		album.Ratings = dna.IntArray{8, 7, 193}
		dna.Log("Id            :", album.Id)
		dna.Log("Title         :", album.Title)
		dna.Log("Artistids     :", album.Artistids)
		dna.Log("Artists       :", album.Artists)
		// dna.Log("Discographies :", album.Discographies)
		// dna.Log("Review        :", album.Review)
		dna.Log("Coverart      :", album.Coverart)
		dna.Log("Duration      :", album.Duration)
		dna.Log("Ratings       :", album.Ratings)
		// dna.Log("Similars      :", album.Similars)
		dna.Log("Genres        :", album.Genres)
		dna.Log("Styles        :", album.Styles)
		dna.Log("Moods         :", album.Moods)
		dna.Log("Themes        :", album.Themes)
		// dna.Log("Songs         :", album.Songs)
		// dna.Log("Releases      :", album.Releases)
		// dna.Log("Awards        :", album.Awards)
		dna.Log("DateReleased  :", dna.Sprintf("%v", album.DateReleased))
		// dna.Log("Credits       :", album.Credits)
		dna.Log("Checktime     :", dna.Sprintf("%v", album.Checktime))

		if album.Discographies != `[{"Id":116253,"Title":"Santana","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0002/104/MI0002104133.jpg?partner=allrovi.com"},{"Id":191745,"Title":"Abraxas","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/862/MI0001862479.jpg?partner=allrovi.com"},{"Id":357213,"Title":"Santana III","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/039/MI0000039691.jpg?partner=allrovi.com"},{"Id":195600,"Title":"Love Devotion Surrender","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/732/MI0001732207.jpg?partner=allrovi.com"},{"Id":650738,"Title":"Caravanserai","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/950/MI0001950906.jpg?partner=allrovi.com"},{"Id":691174,"Title":"Welcome","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0002/110/MI0002110103.jpg?partner=allrovi.com"},{"Id":320315,"Title":"Lotus","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/038/MI0000038850.jpg?partner=allrovi.com"},{"Id":652611,"Title":"Borboletta","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/656/MI0001656513.jpg?partner=allrovi.com"},{"Id":195601,"Title":"Amigos","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0002/424/MI0002424158.jpg?partner=allrovi.com"},{"Id":309055,"Title":"Festival","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0002/455/MI0002455313.jpg?partner=allrovi.com"},{"Id":191747,"Title":"Moonflower","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/677/MI0000677689.jpg?partner=allrovi.com"},{"Id":191626,"Title":"Inner Secrets","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0002/805/MI0002805555.jpg?partner=allrovi.com"},{"Id":650739,"Title":"Marathon","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/511/MI0001511111.jpg?partner=allrovi.com"},{"Id":196513,"Title":"The Swing of Delight","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/495/MI0001495018.jpg?partner=allrovi.com"},{"Id":191748,"Title":"Zebop!","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/039/MI0000039700.jpg?partner=allrovi.com"},{"Id":650740,"Title":"Shangó","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/039/MI0000039693.jpg?partner=allrovi.com"},{"Id":195822,"Title":"Havana Moon","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/038/MI0000038844.jpg?partner=allrovi.com"},{"Id":191625,"Title":"Beyond Appearances","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/038/MI0000038832.jpg?partner=allrovi.com"},{"Id":191746,"Title":"Freedom","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/038/MI0000038840.jpg?partner=allrovi.com"},{"Id":193478,"Title":"Blues for Salvador","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/666/MI0001666207.jpg?partner=allrovi.com"},{"Id":308193,"Title":"Spirits Dancing in the Flesh","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/900/MI0001900971.jpg?partner=allrovi.com"},{"Id":72179,"Title":"Milagro","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/545/MI0001545085.jpg?partner=allrovi.com"},{"Id":104810,"Title":"Sacred Fire: Santana Live in South America","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/541/MI0001541953.jpg?partner=allrovi.com"},{"Id":119213,"Title":"Santana Brothers","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/053/MI0000053980.jpg?partner=allrovi.com"},{"Id":616811,"Title":"Live at the Fillmore 1968","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/126/MI0000126463.jpg?partner=allrovi.com"},{"Id":963904,"Title":"Santana [Direct Source]","Coverart":"http://cps-static.rovicorp.com/3/JPG_170/MI0001/851/MI0001851959.jpg?partner=allrovi.com"},{"Id":229521,"Title":"Shaman","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/366/MI0000366634.jpg?partner=allrovi.com"},{"Id":231840,"Title":"Fillmore Performance: San Francisco 1968","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0000/373/MI0000373425.jpg?partner=allrovi.com"},{"Id":26726,"Title":"Santana: The Collection - Santana/Abraxas/Santana III","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/977/MI0001977184.jpg?partner=allrovi.com"},{"Id":207869,"Title":"All That I Am","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0001/908/MI0001908720.jpg?partner=allrovi.com"},{"Id":2023076,"Title":"Guitar Heaven: The Greatest Guitar Classics of All Time","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0003/001/MI0003001222.jpg?partner=allrovi.com"},{"Id":2341608,"Title":"Shape Shifter","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0003/350/MI0003350654.jpg?partner=allrovi.com"},{"Id":2618166,"Title":"Corazón","Coverart":"http://cps-static.rovicorp.com/3/JPG_75/MI0003/728/MI0003728458.jpg?partner=allrovi.com"}]` {
			panic("Wrong discographies: " + album.Discographies.String())
		}

		if album.Review != `<a href="http://www.allmusic.com/artist/santana-mn0000295756" class="name-link">Santana</a> was still a respected rock veteran in 1999, but it had been years since he had a hit, even if he continued to fare well on the concert circuits. <a href="http://www.allmusic.com/artist/clive-davis-mn0000788138" class="name-link">Clive Davis</a>, the man who had signed <a href="http://www.allmusic.com/artist/santana-mn0000295756" class="name-link">Santana</a> to Columbia in 1968, offered him the opportunity to set up shop at his label, Arista. In the tradition of comebacks and label debuts by veteran artists in the '90s, <a href="http://www.allmusic.com/album/supernatural-mw0000243930" class="album-link">Supernatural</a>, <a href="http://www.allmusic.com/artist/santana-mn0000295756" class="name-link">Santana</a>'s first effort for Arista, is designed as a star-studded event. At first listen, there doesn't seem to be a track that doesn't have a guest star, which brings up the primary problem with the album -- despite several interesting or excellent moments, it never develops a consistent voice that holds the album together. The fault doesn't lay with the guest stars or even with <a href="http://www.allmusic.com/artist/santana-mn0000295756" class="name-link">Santana</a>, who continues to turn in fine performances. There's just a general directionless feeling to the record, enhanced by several songs that seem like excuses for jams, which, truth be told, isn't all that foreign on latter-day <a href="http://www.allmusic.com/artist/santana-mn0000295756" class="name-link">Santana</a> records. Then again, the grooves often play better than the ploys for radio play, but that's not always the case, since <a href="http://www.allmusic.com/artist/lauryn-hill-mn0000113753" class="name-link">Lauryn Hill</a>'s "Do You Like the Way" and <a href="http://www.allmusic.com/artist/the-dust-brothers-mn0000141852" class="name-link">the Dust Brothers</a>-produced, <a href="http://www.allmusic.com/artist/eagle-eye-cherry-mn0000166291" class="name-link">Eagle-Eye Cherry</a>-sung "Wishing It Was" are as captivating as the <a href="http://www.allmusic.com/artist/eric-clapton-mn0000187478" class="name-link">Eric Clapton</a> duet, "The Calling." But that just confirms that <a href="http://www.allmusic.com/album/supernatural-mw0000243930" class="album-link">Supernatural</a> just doesn't have much of a direction, flipping between traditional <a href="http://www.allmusic.com/artist/santana-mn0000295756" class="name-link">Santana</a> numbers and polished contemporary collaborations, with both extremes being equally likely to hit or miss. That doesn't quite constitute a triumph, but the peak moments of <a href="http://www.allmusic.com/album/supernatural-mw0000243930" class="album-link">Supernatural</a> are some of <a href="http://www.allmusic.com/artist/santana-mn0000295756" class="name-link">Santana</a>'s best music of the '90s, which does make it a successful comeback.` {
			panic("Wrong review")
		}

		expectedSongsStr := dna.String(`[{"Id":95440,"Title":"(Da Le) Yaleo","Artists":[{"Id":295756,"Name":"Santana"}],"Composers":[{"Id":0,"Name":""}],"Duration":351},{"Id":1758756,"Title":"Love of My Life","Artists":[{"Id":295756,"Name":"Santana"},{"Id":687561,"Name":"Dave Matthews"}],"Composers":[{"Id":687561,"Name":"Dave Matthews"},{"Id":175196,"Name":"Carlos Santana"}],"Duration":348},{"Id":5248997,"Title":"Put Your Lights On","Artists":[{"Id":295756,"Name":"Santana"},{"Id":155044,"Name":"Everlast"}],"Composers":[{"Id":1013784,"Name":"Erik Schrody"}],"Duration":287},{"Id":5100450,"Title":"Africa Bamba","Artists":[{"Id":295756,"Name":"Santana"}],"Composers":[{"Id":360749,"Name":"Karl Perazzo"},{"Id":175196,"Name":"Carlos Santana"},{"Id":1735694,"Name":"Ismaïla Touré"},{"Id":1286126,"Name":"Sixu Tidiane Touré"}],"Duration":280},{"Id":4867457,"Title":"Smooth","Artists":[{"Id":295756,"Name":"Santana"},{"Id":1009627,"Name":"Rob Thomas"}],"Composers":[{"Id":110666,"Name":"Itaal Shur"},{"Id":1009627,"Name":"Rob Thomas"}],"Duration":296},{"Id":165513,"Title":"Do You Like the Way","Artists":[{"Id":295756,"Name":"Santana"},{"Id":153315,"Name":"Cee Lo Green"},{"Id":113753,"Name":"Lauryn Hill"}],"Composers":[{"Id":1778195,"Name":"L.A. Hill"}],"Duration":352},{"Id":1494945,"Title":"Maria, Maria","Artists":[{"Id":295756,"Name":"Santana"},{"Id":490066,"Name":"The Product G\u0026B"}],"Composers":[{"Id":328303,"Name":"Jerry \"Wonda\" Duplessis"},{"Id":583161,"Name":"Wyclef Jean"},{"Id":360749,"Name":"Karl Perazzo"},{"Id":399694,"Name":"Raul Rekow"},{"Id":175196,"Name":"Carlos Santana"}],"Duration":261},{"Id":5475725,"Title":"Migra","Artists":[{"Id":295756,"Name":"Santana"}],"Composers":[{"Id":744372,"Name":"Tony Lindsay"},{"Id":175196,"Name":"Carlos Santana"},{"Id":1564728,"Name":"Rafi Taha"}],"Duration":324},{"Id":2733928,"Title":"Corazon Espinado","Artists":[{"Id":295756,"Name":"Santana"},{"Id":673621,"Name":"Maná"}],"Composers":[{"Id":795856,"Name":"Fher Olvera"}],"Duration":272},{"Id":2004975,"Title":"Wishing It Was","Artists":[{"Id":295756,"Name":"Santana"},{"Id":166291,"Name":"Eagle-Eye Cherry"}],"Composers":[{"Id":166291,"Name":"Eagle-Eye Cherry"}],"Duration":299},{"Id":8284577,"Title":"El Farol","Artists":[{"Id":295756,"Name":"Santana"}],"Composers":[{"Id":353383,"Name":"KC Porter"},{"Id":175196,"Name":"Carlos Santana"}],"Duration":289},{"Id":5270266,"Title":"Primavera","Artists":[{"Id":295756,"Name":"Santana"}],"Composers":[{"Id":89355,"Name":"Cheín García Alonso"},{"Id":1796803,"Name":"J.B. Eckl"},{"Id":353383,"Name":"KC Porter"},{"Id":92479,"Name":"Kike Santander"}],"Duration":317},{"Id":1102465,"Title":"The Calling","Artists":[{"Id":295756,"Name":"Santana"},{"Id":187478,"Name":"Eric Clapton"}],"Composers":[{"Id":1251331,"Name":"Linda Graham"},{"Id":175196,"Name":"Carlos Santana"},{"Id":188050,"Name":"Freddie Stone"},{"Id":1259629,"Name":"Chester Thompson"}],"Duration":468}]`)
		if album.Songs.Length() != expectedSongsStr.Length() {
			panic("Wrong songs: LENGTHS: " + album.Songs.Length().ToString().String() + " " + expectedSongsStr.Length().ToString().String())
		}

		if album.Releases != `[{"Id":3076285,"Title":"Supernatural","Format":"CD","Year":1999,"Label":"Arista"},{"Id":74660,"Title":"Supernatural","Format":"Cassette","Year":1999,"Label":"Arista"},{"Id":102795,"Title":"Supernatural","Format":"CD","Year":1999,"Label":"Arista"},{"Id":3960316,"Title":"Supernatural","Format":"CD","Year":1999,"Label":"Arista"},{"Id":1714986,"Title":"Supernatural [Bonus Track]","Format":"DVD-Audio","Year":2003,"Label":"Arista"},{"Id":693499,"Title":"Supernatural","Format":"CD","Year":2005,"Label":"Phantom Import Distribution"},{"Id":894585,"Title":"Supernatural","Format":"CD","Year":2006,"Label":"Arista"},{"Id":1703460,"Title":"Supernatural","Format":"CD","Year":2009,"Label":"Arista"},{"Id":1703794,"Title":"Supernatural [Blu-Spec]","Format":"Blu-Ray Disc","Year":2009,"Label":"Sony Music Distribution"},{"Id":1365590,"Title":"Supernatural [Legacy Edition]","Format":"CD","Year":2010,"Label":"Arista / Legacy"}]` {
			panic("wrong release")
		}

		if album.Awards != `[{"Name":"Billboard Albums","Type":"","Awards":[{"Id":243930,"Title":"Supernatural","Year":2010,"Chart":"Top Pop Catalog","Peak":8,"Type":0,"Award":"","Winners":[]},{"Id":243930,"Title":"Supernatural","Year":1999,"Chart":"Latin Pop","Peak":1,"Type":0,"Award":"","Winners":[]},{"Id":243930,"Title":"Supernatural","Year":1999,"Chart":"The Billboard 200","Peak":1,"Type":0,"Award":"","Winners":[]},{"Id":243930,"Title":"Supernatural","Year":1999,"Chart":"Top Canadian Albums","Peak":1,"Type":0,"Award":"","Winners":[]},{"Id":243930,"Title":"Supernatural","Year":1999,"Chart":"Top Internet Albums","Peak":1,"Type":0,"Award":"","Winners":[]},{"Id":243930,"Title":"Supernatural","Year":1999,"Chart":"Top Latin Albums","Peak":1,"Type":0,"Award":"","Winners":[]}]},{"Name":"Billboard Singles","Type":"SINGLE","Awards":[{"Id":5270266,"Title":"Primavera","Year":2001,"Chart":"Latin Tropical/Salsa Airplay","Peak":23,"Type":2,"Award":"","Winners":[]},{"Id":2733928,"Title":"Corazon Espinado","Year":2000,"Chart":"Latin Pop Airplay","Peak":22,"Type":2,"Award":"","Winners":[]},{"Id":1758756,"Title":"Love Of My Life","Year":2000,"Chart":"Adult Top 40","Peak":39,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Adult Top 40","Peak":12,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Canadian Singles Chart","Peak":3,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Hot Dance Music/Maxi-Singles Sales","Peak":1,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Hot Dance Music/Maxi-Singles Sales","Peak":9,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Hot R\u0026B/Hip-Hop Singles \u0026 Tracks","Peak":1,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Latin Pop Airplay","Peak":24,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Latin Tropical/Salsa Airplay","Peak":18,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"The Billboard Hot 100","Peak":1,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Top 40 Mainstream","Peak":2,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":2000,"Chart":"Top 40 Tracks","Peak":2,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":2000,"Chart":"Latin Tropical/Salsa Airplay","Peak":40,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":2000,"Chart":"Top 40 Adult Recurrents","Peak":1,"Type":2,"Award":"","Winners":[]},{"Id":1494945,"Title":"Maria Maria","Year":1999,"Chart":"Rhythmic Top 40","Peak":7,"Type":2,"Award":"","Winners":[]},{"Id":5248997,"Title":"Put Your Lights On","Year":1999,"Chart":"Mainstream Rock","Peak":8,"Type":2,"Award":"","Winners":[]},{"Id":5248997,"Title":"Put Your Lights On","Year":1999,"Chart":"Modern Rock Tracks","Peak":17,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Adult Contemporary","Peak":11,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Adult Top 40","Peak":1,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Canadian Singles Chart","Peak":13,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Mainstream Rock","Peak":10,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Modern Rock Tracks","Peak":24,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Rhythmic Top 40","Peak":23,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"The Billboard Hot 100","Peak":1,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Top 40 Mainstream","Peak":1,"Type":2,"Award":"","Winners":[]},{"Id":4867457,"Title":"Smooth","Year":1999,"Chart":"Top 40 Tracks","Peak":1,"Type":2,"Award":"","Winners":[]}]},{"Name":"Grammy Awards","Type":"ALBUMs \u0026 SONGS","Awards":[{"Id":0,"Title":"Calling","Year":1999,"Chart":"","Peak":0,"Type":2,"Award":"Best Rock Instrumental Performance","Winners":[{"Id":187478,"Name":"Eric Clapton"},{"Id":295756,"Name":"Santana"}]},{"Id":0,"Title":"Farol","Year":1999,"Chart":"","Peak":0,"Type":2,"Award":"Best Pop Instrumental Performance","Winners":[{"Id":295756,"Name":"Santana"}]},{"Id":0,"Title":"Maria, Maria","Year":1999,"Chart":"","Peak":0,"Type":2,"Award":"Best Pop Performance By A Duo Or Group With Vocal","Winners":[{"Id":295756,"Name":"Santana"}]},{"Id":0,"Title":"Put Your Lights On","Year":1999,"Chart":"","Peak":0,"Type":2,"Award":"Best Rock Performance By A Duo Or Group With Vocal","Winners":[{"Id":155044,"Name":"Everlast"},{"Id":295756,"Name":"Santana"}]},{"Id":0,"Title":"Smooth","Year":1999,"Chart":"","Peak":0,"Type":2,"Award":"Best Pop Collaboration with Vocals","Winners":[{"Id":295756,"Name":"Santana"}]},{"Id":0,"Title":"Smooth","Year":1999,"Chart":"","Peak":0,"Type":2,"Award":"Record Of The Year","Winners":[{"Id":811742,"Name":"David Thoener"},{"Id":322483,"Name":"Matt Serletic"},{"Id":1009627,"Name":"Rob Thomas"},{"Id":295756,"Name":"Santana"}]},{"Id":0,"Title":"Smooth","Year":1999,"Chart":"","Peak":0,"Type":2,"Award":"Song Of The Year/New Song Of The Year","Winners":[{"Id":110666,"Name":"Itaal Shur"},{"Id":1009627,"Name":"Rob Thomas"}]},{"Id":0,"Title":"Supernatural","Year":1999,"Chart":"","Peak":0,"Type":1,"Award":"Album Of The Year","Winners":[{"Id":1456657,"Name":"Alex Gonzales"},{"Id":10280,"Name":"Alvaro Villagra"},{"Id":36576,"Name":"Andy Grassi"},{"Id":491079,"Name":"Anton Pukshansky"},{"Id":791803,"Name":"Benny Faccone"},{"Id":1876330,"Name":"Charles Goodan"},{"Id":119820,"Name":"Chris Theis"},{"Id":788138,"Name":"Clive Davis"},{"Id":1804188,"Name":"Comissioner Gordon"},{"Id":120238,"Name":"Dante Ross"},{"Id":534111,"Name":"David Frazer"},{"Id":811742,"Name":"David Thoener"},{"Id":795856,"Name":"Fher Olvera"},{"Id":666955,"Name":"Glen Kolotkin"},{"Id":1794882,"Name":"Hodge, Art"},{"Id":812701,"Name":"Jeff Poe"},{"Id":328303,"Name":"Jerry Duplessis"},{"Id":345005,"Name":"Jim Gaines"},{"Id":348633,"Name":"Jim Scott"},{"Id":223670,"Name":"John Gamble"},{"Id":231715,"Name":"John Karpowich"},{"Id":194341,"Name":"John Seymour"},{"Id":353383,"Name":"K.C. Porter"},{"Id":113753,"Name":"Lauryn Hill"},{"Id":322483,"Name":"Matt Serletic"},{"Id":863827,"Name":"Matty Spindel"},{"Id":893072,"Name":"Mike Couzzi"},{"Id":295756,"Name":"Santana"},{"Id":749898,"Name":"Stephen Harris"},{"Id":1763103,"Name":"Steve Farrone"},{"Id":32105,"Name":"Steve Fontano"},{"Id":5093,"Name":"T-Ray"},{"Id":141852,"Name":"The Dust Brothers"},{"Id":608716,"Name":"Tom Lord-Alge"},{"Id":15712,"Name":"Tony Prendatt"},{"Id":196547,"Name":"Warren Riker"},{"Id":583161,"Name":"Wyclef Jean"}]},{"Id":0,"Title":"Supernatural","Year":1999,"Chart":"","Peak":0,"Type":1,"Award":"Best Rock Album","Winners":[{"Id":788138,"Name":"Clive Davis"},{"Id":295756,"Name":"Santana"},{"Id":32105,"Name":"Steve Fontano"}]}]}]` {
			panic("Wrong award")
		}

		if album.Credits != `[{"Id":1951569,"Artist":"Gonzales Alex","Job":"Vocals (Background)"},{"Id":89355,"Artist":"Cheín García Alonso","Job":"Composer"},{"Id":614155,"Artist":"Al Anderson","Job":"Guitar (Rhythm)"},{"Id":930639,"Artist":"Tom Barney","Job":"Bass"},{"Id":1879286,"Artist":"Kobie Brown","Job":"Programming"},{"Id":166291,"Artist":"Eagle-Eye Cherry","Job":"Composer, Featured Artist, Performer, Primary Artist"},{"Id":187478,"Artist":"Eric Clapton","Job":"Featured Artist, Guest Artist, Performer, Primary Artist"},{"Id":269829,"Artist":"Jeremy Cohen","Job":"Violin"},{"Id":1804188,"Artist":"Comissioner Gordon","Job":"Engineer"},{"Id":893072,"Artist":"Mike Couzzi","Job":"Engineer"},{"Id":812178,"Artist":"Jeff Cressman","Job":"Trombone"},{"Id":223844,"Artist":"David Crockett","Job":"Drums, Percussion"},{"Id":69174,"Artist":"Joseph Daley","Job":"Tuba"},{"Id":788138,"Artist":"Clive Davis","Job":"Producer"},{"Id":191224,"Artist":"Francis Dunnery","Job":"Guitar (Rhythm)"},{"Id":328303,"Artist":"Jerry \"Wonda\" Duplessis","Job":"Composer, Producer"},{"Id":141852,"Artist":"The Dust Brothers","Job":"Producer"},{"Id":1796803,"Artist":"J.B. Eckl","Job":"Composer"},{"Id":155044,"Artist":"Everlast","Job":"Featured Artist, Guest Artist, Performer, Primary Artist"},{"Id":791803,"Artist":"Benny Faccone","Job":"Engineer, Mixing"},{"Id":1763103,"Artist":"Steve Farrone","Job":"Engineer"},{"Id":144560,"Artist":"Fher","Job":"Vocals"},{"Id":1865039,"Artist":"Abel Figueroa","Job":"Trombone"},{"Id":146976,"Artist":"Ramon Flores","Job":"Trombone"},{"Id":32105,"Artist":"Steve Fontano","Job":"Engineer"},{"Id":534111,"Artist":"David Frazer","Job":"Engineer"},{"Id":1784533,"Artist":"G\u0026B Product","Job":"Vocals"},{"Id":345005,"Artist":"Jim Gaines","Job":"Engineer"},{"Id":223670,"Artist":"John Gamble","Job":"Engineer"},{"Id":840069,"Artist":"Pete Ganbarg","Job":"A\u0026R"},{"Id":261838,"Artist":"Earl Gardner","Job":"Flugelhorn, Trumpet"},{"Id":908034,"Artist":"Mic Gillette","Job":"Trumpet"},{"Id":1456657,"Artist":"Alex Gonzales","Job":"Producer"},{"Id":1876330,"Artist":"Charles Goodan","Job":"Producer"},{"Id":1251331,"Artist":"Linda Graham","Job":"Composer"},{"Id":36576,"Artist":"Andy Grassi","Job":"Engineer"},{"Id":153315,"Artist":"Cee Lo Green","Job":"Featured Artist, Guest Artist, Primary Artist, Vocals"},{"Id":749898,"Artist":"Stephen Harris","Job":"Producer"},{"Id":205290,"Artist":"Joseph Hébert","Job":"Cello"},{"Id":219569,"Artist":"Horacio \"El Negro\" Hernández","Job":"Drums"},{"Id":1778195,"Artist":"L.A. Hill","Job":"Composer"},{"Id":113753,"Artist":"Lauryn Hill","Job":"Featured Artist, Guest Artist, Performer, Primary Artist, Producer"},{"Id":1794882,"Artist":"Hodge, Art","Job":"Producer"},{"Id":283632,"Artist":"Loris Holland","Job":"Keyboards"},{"Id":583161,"Artist":"Wyclef Jean","Job":"Composer, Guest Artist, Producer"},{"Id":231715,"Artist":"John Karpowich","Job":"Engineer"},{"Id":666955,"Artist":"Glen Kolotkin","Job":"Engineer"},{"Id":744372,"Artist":"Tony Lindsay","Job":"Composer, Vocals (Background)"},{"Id":608716,"Artist":"Tom Lord-Alge","Job":"Engineer"},{"Id":673621,"Artist":"Maná","Job":"Featured Artist, Guest Artist, Performer, Primary Artist"},{"Id":1532629,"Artist":"René Martínez","Job":"Guitar"},{"Id":687561,"Artist":"Dave Matthews","Job":"Composer, Featured Artist, Guest Artist, Performer, Primary Artist"},{"Id":102710,"Artist":"Marvin McFadden","Job":"Trumpet"},{"Id":834108,"Artist":"Sarah McLachlan","Job":"Guest Artist"},{"Id":836412,"Artist":"Julius Melendez","Job":"Trumpet"},{"Id":518098,"Artist":"Brian Montgomery","Job":"Monitors"},{"Id":2286594,"Artist":"Patrick Murray","Job":"House Sound"},{"Id":2158343,"Artist":"Nico Nibbering","Job":"Keyboards"},{"Id":795856,"Artist":"Fher Olvera","Job":"Composer, Producer"},{"Id":626597,"Artist":"William Oritz","Job":"Trumpet"},{"Id":426564,"Artist":"Ozomatli","Job":"Guest Artist"},{"Id":360749,"Artist":"Karl Perazzo","Job":"Composer, Percussion, Vocals (Background)"},{"Id":812701,"Artist":"Jeff Poe","Job":"Engineer"},{"Id":86349,"Artist":"Che Pope","Job":"Programming"},{"Id":353383,"Artist":"KC Porter","Job":"Accordion, Composer, Producer, Programming, Vocals"},{"Id":15712,"Artist":"Tony Prendatt","Job":"Engineer"},{"Id":490066,"Artist":"The Product G\u0026B","Job":"Featured Artist, Guest Artist"},{"Id":491079,"Artist":"Anton Pukshansky","Job":"Engineer"},{"Id":499370,"Artist":"Lenesha Randolph","Job":"Vocals (Background)"},{"Id":399694,"Artist":"Raul Rekow","Job":"Composer, Congas"},{"Id":51133,"Artist":"Benny Rietveld","Job":"Bass"},{"Id":196547,"Artist":"Warren Riker","Job":"Engineer"},{"Id":120238,"Artist":"Dante Ross","Job":"Producer"},{"Id":933472,"Artist":"Alberto Salas","Job":"Keyboards"},{"Id":295756,"Artist":"Santana","Job":"Primary Artist"},{"Id":175196,"Artist":"Carlos Santana","Job":"Composer, Guitar, Performer, Sleigh Bells, Vocals, Vocals (Background)"},{"Id":92479,"Artist":"Kike Santander","Job":"Composer"},{"Id":1013784,"Artist":"Erik Schrody","Job":"Composer"},{"Id":348633,"Artist":"Jim Scott","Job":"Engineer"},{"Id":679573,"Artist":"Danny Seidenberg","Job":"Viola"},{"Id":322483,"Artist":"Matt Serletic","Job":"Producer"},{"Id":194341,"Artist":"John Seymour","Job":"Engineer"},{"Id":250435,"Artist":"Wayne Shorter","Job":"Guest Artist"},{"Id":110666,"Artist":"Itaal Shur","Job":"Composer"},{"Id":863827,"Artist":"Matty Spindel","Job":"Engineer"},{"Id":188050,"Artist":"Freddie Stone","Job":"Composer"},{"Id":1991209,"Artist":"Angus Sutherland","Job":"Guitar"},{"Id":1564728,"Artist":"Rafi Taha","Job":"Composer"},{"Id":119820,"Artist":"Chris Theis","Job":"Engineer"},{"Id":811742,"Artist":"David Thoener","Job":"Engineer, Mixing"},{"Id":1009627,"Artist":"Rob Thomas","Job":"Composer, Featured Artist, Guest Artist, Performer, Primary Artist, Vocals"},{"Id":1259629,"Artist":"Chester Thompson","Job":"Composer, Keyboards"},{"Id":1735694,"Artist":"Ismaïla Touré","Job":"Composer"},{"Id":1286126,"Artist":"Sixu Tidiane Touré","Job":"Composer"},{"Id":1374963,"Artist":"Steve Touré","Job":"Trombone"},{"Id":5093,"Artist":"T-Ray","Job":"Engineer"},{"Id":8648,"Artist":"Sergio Vallin","Job":"Guitar (Rhythm)"},{"Id":10280,"Artist":"Alvaro Villagra","Job":"Engineer"},{"Id":1878061,"Artist":"Danny Wolensky","Job":"Flute, Saxophone"}]` {
			panic("Wrong credits")
		}

		if album.Similars.Length() <= 0 {
			panic("Wrong silimars")
		}

	}
	// Output:
	// Id            : 243930
	// Title         : Supernatural
	// Artistids     : dna.IntArray{295756}
	// Artists       : dna.StringArray{"Santana"}
	// Coverart      : http://cps-static.rovicorp.com/3/JPG_250/MI0001/839/MI0001839570.jpg?partner=allrovi.com
	// Duration      : 4144
	// Ratings       : dna.IntArray{8, 7, 193}
	// Genres        : [{"Id":2613,"Name":"Pop/Rock"}]
	// Styles        : [{"Id":4488,"Name":"Adult Alternative Pop/Rock"},{"Id":4459,"Name":"Album Rock"},{"Id":12230,"Name":"Alternative/Indie Rock"},{"Id":2468,"Name":"Blues-Rock"},{"Id":2636,"Name":"Hard Rock"},{"Id":11908,"Name":"Latin Rock"},{"Id":4443,"Name":"Contemporary Pop/Rock"}]
	// Moods         : [{"Id":1116,"Name":"Summery"},{"Id":949,"Name":"Brash"},{"Id":703,"Name":"Celebratory"},{"Id":704,"Name":"Cheerful"},{"Id":981,"Name":"Earthy"},{"Id":997,"Name":"Exuberant"},{"Id":1029,"Name":"Joyous"},{"Id":1057,"Name":"Organic"},{"Id":1062,"Name":"Passionate"},{"Id":759,"Name":"Rousing"},{"Id":764,"Name":"Sensual"},{"Id":1105,"Name":"Spiritual"}]
	// Themes        : [{"Id":6289,"Name":"At the Office"},{"Id":6330,"Name":"Day Driving"},{"Id":6297,"Name":"Hanging Out"},{"Id":4281,"Name":"Summer"},{"Id":6308,"Name":"TGIF"}]
	// DateReleased  : 1999-06-15 00:00:00 +0000 UTC
	// Checktime     : 2013-11-21 00:00:00 +0000 UTC
}
