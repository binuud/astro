package astro

import astroV1 "github.com/binuud/astro/gen/go/v1/astro"

const (
	PLANET_ASC     = "asc"
	PLANET_SUN     = "sun"
	PLANET_MOON    = "moon"
	PLANET_MERCURY = "mercury"
	PLANET_VENUS   = "venus"
	PLANET_EARTH   = "earth"
	PLANET_MARS    = "mars"
	PLANET_JUPITER = "jupiter"
	PLANET_SATURN  = "saturn"
	PLANET_URANUS  = "uranus"
	PLANET_NEPTUNE = "neptune"
	PLANET_PLUTO   = "pluto"
	PLANET_RAHU    = "rahu"
	PLANET_KETU    = "ketu"
	PLANET_MAANDI  = "maandi"
	PLANET_GULIKA  = "gulkia"

	KRISHNA_PAKSHA = 2
	SHUKRA_PAKSHA  = 1

	HOUR_MINUTES        = 60
	QUARTER_DAY_MINUTES = 360

	SECONDS_IN_YEAR = 31556952

	RASI_DEGREES   = 30
	STAR_DEGREES   = 13.33333
	PAADA_DEGREES  = 3.333
	TITHI_DEGREES  = 12
	YOGA_DEGREES   = 13.33333
	KARANA_DEGREES = 6

	TOTAL_RASIS              = 12
	TOTAL_STARS              = 27
	TOTAL_STARS_WITH_ABHIJIT = 28
	TOTAL_PAADAS             = 108
	TOTAL_TITHIS             = 30
	TOTAL_YOGAS              = 27
	TOTAL_KARANAS            = 60

	HARMONIC_HORA         = 2  // BiCyclical Hora - Pariviti Dwaya
	HARMONIC_NAVAMSA      = 9  // Navamsa
	HARMONIC_DASAMSA      = 10 // Dasamsa - coming wrong
	HARMONIC_NAKSATHRAMSA = 27
	HARMONIC_SASHTYAMSA   = 60 // Coming wrong

	STAR_VAINASIKA = 22

	COMBUSTION_MOON    = 12
	COMBUSTION_MARS    = 17
	COMBUSTION_MERCURY = 14
	COMBUSTION_JUPITER = 11
	COMBUSTION_VENUS   = 10
	COMBUSTION_SATURN  = 15
)

var KARANA_NAMES = []string{"", "Kintughna",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Bava", "Baalava", "Kaulava", "Taitila", "Gara", "Vanija", "Vishti",
	"Shakuni", "Chatushpada", "Naga"}

var TITHI_NAMES = []string{"", "Prathama", "Dwitiya", "Tritiya", "Chaturthi", "Panchami",
	"Shasti", "Saptami", "Ashtami", "Navami", "Dashami",
	"Ekadashi", "Dwadashi", "Thrayodashi", "Chaturdashi", "Purnima",
	"Prathama", "Dwitiya", "Tritiya", "Chaturthi", "Panchami",
	"Shasti", "Saptami", "Ashtami", "Navami", "Dashami",
	"Ekadashi", "Dwadashi", "Thrayodashi", "Chaturdashi", "Amavasya"}

var STAR_NAMES = []string{"", "Aswini", "Barani", "Krittika",
	"Rohini", "Mrigasira", "Arudra",
	"Punarvasu", "Pushyami", "Aslesha",
	"Makha", "Poorvaphalguni", "Uttaraphalguni",
	"Hasta", "Chitra", "Swati",
	"Vishaka", "Anuradha", "Jyeshta",
	"Moola", "Poorvashada", "Uttarashada",
	"Sravana", "Dhanishta", "Satabhisha",
	"Poorvabhadra", "Uttarabadra", "Revathi"}

var STAR_NAMES_WITH_ABHIJIT = []string{"", "Aswini", "Barani", "Krittika",
	"Rohini", "Mrigasira", "Arudra",
	"Punarvasu", "Pushyami", "Aslesha",
	"Makha", "Poorvaphalguni", "Uttaraphalguni",
	"Hasta", "Chitra", "Swati",
	"Vishaka", "Anuradha", "Jyeshta",
	"Moola", "Poorvashada", "Uttarashada", "Abhijit",
	"Sravana", "Dhanishta", "Satabhisha",
	"Poorvabhadra", "Uttarabadra", "Revathi"}

var RASI_NAMES = []string{"", "Mesham", "Rishabam", "Mithunam", "Karkataka", "Simha", "Kanya", "Thula", "Vrishika", "Dhanus", "Makara", "Kumbha", "Meena"}
var ENGLISH_RASI_NAMES = []string{"", "Aries", "Taurus", "Gemini", "Cancer", "Leo", "Virgo", "Libra", "Scorpio", "Sagittarius", "Capricorn", "Aquarius", "Pisces"}

var RASI_OWNERS = []int{-1, int(astroV1.AstroConstants_Mars), int(astroV1.AstroConstants_Venus), int(astroV1.AstroConstants_Mercury),
	int(astroV1.AstroConstants_Moon), int(astroV1.AstroConstants_Sun), int(astroV1.AstroConstants_Mercury),
	int(astroV1.AstroConstants_Venus), int(astroV1.AstroConstants_Mars), int(astroV1.AstroConstants_Jupiter),
	int(astroV1.AstroConstants_Saturn), int(astroV1.AstroConstants_Saturn), int(astroV1.AstroConstants_Jupiter)}

var PAKSHA_NAMES = []string{"", "Shukra Paksha", "Krishna Paksha"}

var YOGA_NAMES = []string{"", "Vishkumbha", "Preeti", "Ayushman", "Saubhagya", "Shobhana",
	"Atiganda", "Sukarma", "Dhriti", "Shoola", "Ganda",
	"Vriddhi", "Dhruva", "Vyaghata", "Harshana", "Vajra",
	"Siddhi", "Vyatipata", "Variyana", "Parigha", "Shiva",
	"Siddha", "Sadhya", "Shubha", "Shukla", "Brahma",
	"Indra", "Vaidhriti"}

var YOGA_BAD = []string{"", "Vyaghata", "Vajra", "Parigha", "Vyatipata", "Vaidhriti", "Ganda", "Atiganda", "Shoola", "Vishkumbha"}

// # "Purvabhadra", "Uttarabhadra",
// var STAR_LIST = []string {"Ashvini", "Bharani", "Krttika", "Rohini", "Mrigashirsha", "Ardra", "Punarvasu", "Pushya", "Ashlesha",
// "Magha", "Uttaraphalguni", "Purvaphalguni", "Hasta", "Citra", "Vishakha", "Anuradha","Jyeshtha",
// "Mula","Purvashadha","Uttarashadha", "Shravishtha", "Dhanishtha", "Purvabhadra", "Uttarabhadra", "Revati"}
// public static $STARS = array();

var STAR_YOGA = []string{"", "Ananda", "Kaaldanda", " Dhumra", "Dhata", "Saumya", "Dhwansha", "Ketu",
	"Srivatsa", "Vajra", "Mudgara", "Chhatra", "Mitra", "Manas", "Padma",
	"Lumba", "Utpaata", "Mrityu", "Kaana", "Siddhi", "Subha", "Amrita",
	"Mushala", "Gada", "Matanga", "Raksha", "Chara", "Susthira", "Pravardhamama"}

var WEEKDAYS = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
