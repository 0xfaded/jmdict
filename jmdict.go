package jmdict

import (
	"encoding/xml"
	"io"
)

type Position string

const (
	AdjI              Position = "adj-i"     // adjective (keiyoushi)
	AdjKu                      = "adj-ku"    // `ku' adjective (archaic)
	AdjNa                      = "adj-na"    // adjectival nouns or quasi-adjectives (keiyodoshi)
	AdjNari                    = "adj-nari"  // archaic/formal form of na-adjective
	AdjNo                      = "adj-no"    // nouns which may take the genitive case particle `no'
	AdjPreNoun                 = "adj-pn"    // pre-noun adjectival (rentaishi)
	AdjShiku                   = "adj-shiku" // `shiku' adjective (archaic)
	AdjTaru                    = "adj-t"     // `taru' adjective
	AdjFunctional              = "adj-f"     // noun or verb acting prenominally
	Adv                        = "adv"       // adverb (fukushi)
	AdvTo                      = "adv-to"    // adverb taking the `to' particle
	Aux                        = "aux"       // auxiliary
	AuxAdj                     = "aux-adj"   // auxiliary adjective
	AuxVerb                    = "aux-v"     // auxiliary verb
	Conj                       = "conj"      // conjunction
	Counter                    = "ctr"       // counter
	Expression                 = "exp"       // Expressions (phrases, clauses, etc.)
	Interjection               = "int"       // interjection (kandoushi)
	Noun                       = "n"         // noun (common) (futsuumeishi)
	NounAdv                    = "n-adv"     // adverbial noun (fukushitekimeishi)
	NounPropper                = "n-pr"      // proper noun
	NounPrefix                 = "n-pref"    // noun, used as a prefix
	NounSuffix                 = "n-suf"     // noun, used as a suffix
	NounTemporal               = "n-t"       // noun (temporal) (jisoumeishi)
	Numeric                    = "num"       // numeric
	Pronoun                    = "pn"        // pronoun
	Prefix                     = "pref"      // prefix
	Particle                   = "prt"       // particle
	Suffix                     = "suf"       // suffix
	Verb1                      = "v1"        // Ichidan verb
	Verb2as                    = "v2a-s"     // Nidan verb with 'u' ending (archaic)
	Verb2bk                    = "v2b-k"     // Nidan verb (upper class) with `bu' ending (archaic)
	Verb2ds                    = "v2d-s"     // Nidan verb (lower class) with `dzu' ending (archaic)
	Verb2gk                    = "v2g-k"     // Nidan verb (upper class) with `gu' ending (archaic)
	Verb2gs                    = "v2g-s"     // Nidan verb (lower class) with `gu' ending (archaic)
	Verb2hk                    = "v2h-k"     // Nidan verb (upper class) with `hu/fu' ending (archaic)
	Verb2hs                    = "v2h-s"     // Nidan verb (lower class) with `hu/fu' ending (archaic)
	Verb2kk                    = "v2k-k"     // Nidan verb (upper class) with `ku' ending (archaic)
	Verb2ks                    = "v2k-s"     // Nidan verb (lower class) with `ku' ending (archaic)
	Verb2ms                    = "v2m-s"     // Nidan verb (lower class) with `mu' ending (archaic)
	Verb2ns                    = "v2n-s"     // Nidan verb (lower class) with `nu' ending (archaic)
	Verb2rk                    = "v2r-k"     // Nidan verb (upper class) with `ru' ending (archaic)
	Verb2rs                    = "v2r-s"     // Nidan verb (lower class) with `ru' ending (archaic)
	Verb2ss                    = "v2s-s"     // Nidan verb (lower class) with `su' ending (archaic)
	Verb2tk                    = "v2t-k"     // Nidan verb (upper class) with `tsu' ending (archaic)
	Verb2ts                    = "v2t-s"     // Nidan verb (lower class) with `tsu' ending (archaic)
	Verb2ws                    = "v2w-s"     // Nidan verb (lower class) with `u' ending and `we' conjugation (archaic)
	Verb2yk                    = "v2y-k"     // Nidan verb (upper class) with `yu' ending (archaic)
	Verb2ys                    = "v2y-s"     // Nidan verb (lower class) with `yu' ending (archaic)
	Verb2zs                    = "v2z-s"     // Nidan verb (lower class) with `zu' ending (archaic)
	Verb4b                     = "v4b"       // Yodan verb with `bu' ending (archaic)
	Verb4h                     = "v4h"       // Yodan verb with `hu/fu' ending (archaic)
	Verb4k                     = "v4k"       // Yodan verb with `ku' ending (archaic)
	Verb4r                     = "v4r"       // Yodan verb with `ru' ending (archaic)
	Verb4s                     = "v4s"       // Yodan verb with `su' ending (archaic)
	Verb4t                     = "v4t"       // Yodan verb with `tsu' ending (archaic)
	Verb5aru                   = "v5aru"     // Godan verb - -aru special class
	Verb5b                     = "v5b"       // Godan verb with `bu' ending
	Verb5g                     = "v5g"       // Godan verb with `gu' ending
	Verb5k                     = "v5k"       // Godan verb with `ku' ending
	Verb5ks                    = "v5k-s"     // Godan verb - Iku/Yuku special class
	Verb5m                     = "v5m"       // Godan verb with `mu' ending
	Verb5n                     = "v5n"       // Godan verb with `nu' ending
	Verb5r                     = "v5r"       // Godan verb with `ru' ending
	Verb5ri                    = "v5r-i"     // Godan verb with `ru' ending (irregular verb)
	Verb5s                     = "v5s"       // Godan verb with `su' ending
	Verb5t                     = "v5t"       // Godan verb with `tsu' ending
	Verb5u                     = "v5u"       // Godan verb with `u' ending
	Verb5us                    = "v5u-s"     // Godan verb with `u' ending (special class)
	VerbIntransitive           = "vi"        // intransitive verb
	VerbKuru                   = "vk"        // Kuru verb - special class
	VerbNu                     = "vn"        // irregular nu verb
	VerbRu                     = "vr"        // irregular ru verb, plain form ends with -ri
	VerbSuru                   = "vs"        // noun or participle which takes the aux. verb suru
	VerbSu                     = "vs-c"      // su verb - precursor to the modern suru
	VerbIrregularSuru          = "vs-i"      // suru verb - irregular
	VerbSuruSpecial            = "vs-s"      // suru verb - special class
	VerbTransitive             = "vt"        // transitive verb
	VerbZuru                   = "vz"        // Ichidan verb - zuru verb (alternative form of -jiru verbs)
)

type Field string

const (
	Anatomical   Field = "anat"    // anatomical term
	Architecture       = "archit"  // architecture term
	Astronomy          = "astron"  // astronomy, etc. term
	Baseball           = "baseb"   // baseball term
	Biology            = "biol"    // biology term
	Botany             = "bot"     // botany term
	buddhist           = "Buddh"   // Buddhist term
	Business           = "bus"     // business term
	Chemistry          = "chem"    // chemistry term
	Computer           = "comp"    // computer terminology
	Economics          = "econ"    // economics term
	Engineering        = "engr"    // engineering term
	Finance            = "finc"    // finance term
	Food               = "food"    // food term
	Geology            = "geol"    // geology, etc. term
	Geometry           = "geom"    // geometry term
	Law                = "law"     // law, etc. term
	Linguistics        = "ling"    // linguistics terminology
	Martial            = "MA"      // martial arts term
	Mathematics        = "math"    // mathematics
	Medicine           = "med"     // medicine, etc. term
	Military           = "mil"     // military
	Music              = "music"   // music term
	Physics            = "physics" // physics terminology
	shinto             = "Shinto"  // Shinto term
	Sports             = "sports"  // sports term
	Sumo               = "sumo"    // sumo term
	Zoology            = "zool"    // zoology term
)

type Misc string

const (
	Abbreviation   Misc = "abbr"    // abbreviation
	Archaism            = "arch"    // archaism
	ChildLanguage       = "chn"     // children's language
	Colloquialism       = "col"     // colloquialism
	Derogatory          = "derog"   // derogatory
	Familiar            = "fam"     // familiar language
	FemaleLanguage      = "fem"     // female term or language
	Honorific           = "hon"     // honorific or respectful (sonkeigo) language
	Humble              = "hum"     // humble (kenjougo) language
	Idiomatic           = "id"      // idiomatic expression
	Jocular             = "joc"     // jocular, humorous term
	MaleLanguage        = "male"    // male term or language
	Manga               = "m-sl"    // manga slang
	Obsolete            = "obs"     // obsolete term
	Obscure             = "obsc"    // obscure term
	Onomatopoeic        = "on-mim"  // onomatopoeic or mimetic word
	Poetical            = "poet"    // poetical term
	Polite              = "pol"     // polite (teineigo) language
	Proverb             = "proverb" // proverb
	Rare                = "rare"    // rare
	Sensitive           = "sens"    // sensitive
	Slang               = "sl"      // slang
	KanaAlone           = "uk"      // word usually written using kana alone
	Vulgar              = "vulg"    // vulgar expression or word
	XRated              = "X"       // rude or X-rated term (not displayed in educational software)
)

type Dialect string

const (
	HokkaidoBen Dialect = "hob"  // Hokkaido-ben
	KansaiBen           = "ksb"  // Kansai-ben
	KantouBen           = "ktb"  // Kantou-ben
	KyotoBen            = "kyb"  // Kyoto-ben
	KyuushuuBen         = "kyu"  // Kyuushuu-ben
	NaganoBen           = "nab"  // Nagano-ben
	OsakaBen            = "osb"  // Osaka-ben
	RyuukyuuBen         = "rkb"  // Ryuukyuu-ben
	TouhokuBen          = "thb"  // Touhoku-ben
	TosaBen             = "tsb"  // Tosa-ben
	TsugaruBen          = "tsug" // Tsugaru-ben
)

// This is a coded information field related specifically to the orthography of
// the keb or reb, and will typically indicate some unusual aspect, such as
// okurigana irregularity.
type Orthography string

const (
	Ateji                   Orthography = "ateji" // ateji (phonetic) reading
	Gikun                               = "gikun" // gikun (meaning as reading) or jukujikun (special kanji reading)
	IrregularKanji                      = "iK"    // word containing irregular kanji usage
	IrregularKana                       = "ik"    // word containing irregular kana usage
	IrregularOkurigana                  = "io"    // irregular okurigana usage
	OutdatedKanji                       = "oK"    // out-dated or obsolete kanji usage
	OutdatedKana                        = "ok"    // out-dated or obsolete kana usage
	OutdatedOrIrregularKana             = "oik"   // "old or irregular kana form
	KanjiAlone                          = "uK"    // word usually written using kanji alone
)

type JMDict struct {
	XMLName xml.Name `xml:"JMdict"`
	Entries []Entry  `xml:"entry"`
}

// Entries consist of kanji elements, reading elements,
// general information and sense elements. Each entry must have at
// least one reading element and one sense element. Others are optional.
type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Id      EntSeq   `xml:"ent_seq"`
	Kanji   []KEle   `xml:"k_ele"`
	Reading []REle   `xml:"r_ele"`
	Info    Info     `xml:"info"`
	Sense   []Sense  `xml:"sense"`
}

// A unique numeric sequence number for each entry
type EntSeq uint64

// The kanji element, or in its absence, the reading element, is
// the defining component of each entry.
// The overwhelming majority of entries will have a single kanji
// element associated with a word in Japanese. Where there are
// multiple kanji elements within an entry, they will be orthographical
// variants of the same word, either using variations in okurigana, or
// alternative and equivalent kanji. Common "mis-spellings" may be
// included, provided they are associated with appropriate information
// fields. Synonyms are not included; they may be indicated in the
// cross-reference field associated with the sense element.
type KEle struct {
	XMLName  xml.Name      `xml:"k_ele"`
	Phrase   Keb           `xml:"keb"`
	Info     []Orthography `xml:"ke_inf"`
	Priority []KePriority  `xml:"ke_pri"`
}

// This element will contain a word or short phrase in Japanese
// which is written using at least one non-kana character (usually kanji,
// but can be other characters). The valid characters are
// kanji, kana, related characters such as chouon and kurikaeshi, and
// in exceptional cases, letters from other alphabets.
type Keb string

type PriorityCode string

const (
	Newspaper  PriorityCode = "news" // news1/news2
	BunruiShuu              = "ichi" // ichi1/ichi2
	LoanWord                = "gai"  // gai1/gai2
	Special                 = "spec" // spec1/spec2
	Frequency               = "nf"   // nfxx
)

// This and the equivalent re_pri field are provided to record
// information about the relative priority of the entry,  and consist
// of codes indicating the word appears in various references which
// can be taken as an indication of the frequency with which the word
// is used. This field is intended for use either by applications which
// want to concentrate on entries of  a particular priority, or to
// generate subset files.
// The current values in this field are:
// - news1/2: appears in the "wordfreq" file compiled by Alexandre Girardi
// from the Mainichi Shimbun. (See the Monash ftp archive for a copy.)
// Words in the first 12,000 in that file are marked "news1" and words
// in the second 12,000 are marked "news2".
// - ichi1/2: appears in the "Ichimango goi bunruishuu", Senmon Kyouiku
// Publishing, Tokyo, 1998.  (The entries marked "ichi2" were
// demoted from ichi1 because they were observed to have low
// frequencies in the WWW and newspapers.)
// - spec1 and spec2: a small number of words use this marker when they
// are detected as being common, but are not included in other lists.
// - gai1/2: common loanwords, based on the wordfreq file.
// - nfxx: this is an indicator of frequency-of-use ranking in the
// wordfreq file. "xx" is the number of the set of 500 words in which
// the entry can be found, with "01" assigned to the first 500, "02"
// to the second, and so on. (The entries with news1, ichi1, spec1 and
// gai1 values are marked with a "(P)" in the EDICT and EDICT2
// files.)
// The reason both the kanji and reading elements are tagged is because
// on occasions a priority is only associated with a particular
// kanji/reading pair.
type Priority struct {
	// String representation of Priority which other fields
	// will be parsed from after unmarshalling
	Raw string `xml:",chardata"`

	Code PriorityCode `xml:"-"`
	Rank int          `xml:"-"`
}

type KePriority struct {
	XMLName xml.Name `xml:"ke_pri"`
	Priority
}

type RePriority struct {
	XMLName xml.Name `xml:"re_pri"`
	Priority
}

// The reading element typically contains the valid readings
// of the word(s) in the kanji element using modern kanadzukai.
// Where there are multiple reading elements, they will typically be
// alternative readings of the kanji element. In the absence of a
// kanji element, i.e. in the case of a word or phrase written
// entirely in kana, these elements will define the entry.
type REle struct {
	XMLName xml.Name `xml:"r_ele"`
	Phrase  Reb      `xml:"reb"`

	// Indicates that the reb, while associated with the keb, cannot be
	// regarded as a true reading of the kanji. It is typically used for
	// words such as foreign place names, gairaigo which can be in kanji or
	// katakana, etc. A non-nil value represents true
	ImproperReading *string `xml:"re_nokanji"`

	Restrict    []ReRestr     `xml:"re_restr"`
	Orthography []Orthography `xml:"re_inf"`
	Priority    []RePriority  `xml:"re_pri"`
}

// this element content is restricted to kana and related
// characters such as chouon and kurikaeshi. Kana usage will be
// consistent between the keb and reb elements; e.g. if the keb
// contains katakana, so too will the reb.
type Reb string

// This element is used to indicate when the reading only applies
// to a subset of the keb elements in the entry. In its absence, all
// readings apply to all kanji elements. The contents of this element
// must exactly match those of one of the keb elements.
type ReRestr string

type Info struct {
	XMLName xml.Name `xml:"info"`
	Links   []Links  `xml:"links"`
	Bibl    []Bibl   `xml:"bibl"`
	Etym    []string `xml:"etym"`
	Audit   []Audit  `xml:"audit"`
}

// Bibliographic information about the entry. The bib_tag will a
// coded reference to an entry in an external bibliographic database.
type Bibl struct {
	XMLName xml.Name `xml:"bibl"`
	BibTag  string   `xml:"bib_tag"`
	BibTxt  string   `xml:"bib_txt"`
}

// This element holds details of linking information to
// entries in other electronic repositories. The link_tag will be
// coded to indicate the type of link (text, image, sound), the
// link_desc will provided a textual label for the link, and the
// link_uri contains the actual URI.
type Links struct {
	XMLName  xml.Name `xml:"links"`
	LinkTag  string   `xml:"link_tag"`
	LinkDesc string   `xml:"link_desc"`
	LinkUri  string   `xml:"link_uri"`
}

// The audit element will contain the date and other information
// about updates to the entry. Can be used to record the source of
type Audit struct {
	XMLName xml.Name `xml:"audit"`
	UpdDate string   `xml:"upd_date"`
	UpdDetl string   `xml:"upd_detl"`
}

// The sense element will record the translational equivalent
// of the Japanese word, plus other related information. Where there
// are several distinctly different meanings of the word, multiple
// sense elements will be employed.
type Sense struct {
	XMLName xml.Name `xml:"sense"`

	// Indicate that the sense is restricted to listed Kanji usages
	KanjiRestrict []string `xml:"stagk"`

	// Indicate that the sense is restricted to listed Reading usages
	ReadingRestrict []string `xml:"stagr"`

	// Part-of-speech information about the entry/sense. Should use
	// appropriate entity codes. In general where there are multiple senses
	// in an entry, the part-of-speech of an earlier sense will apply to
	// later senses unless there is a new part-of-speech indicated.
	Position []Position `xml:"pos"`

	// This element is used to indicate a cross-reference to another
	// entry with a similar or related meaning or sense. The content of
	// this element is typically a keb or reb element in another entry. In some
	// cases a keb will be followed by a reb and/or a sense number to provide
	// a precise target for the cross-reference. Where this happens, a JIS
	// "centre-dot" (0x2126) is placed between the components of the
	// cross-reference.
	Xref []string `xml:"xref"`

	// This element is used to indicate another entry which is an
	// antonym of the current entry/sense. The content of this element
	// must exactly match that of a keb or reb element in another entry.
	Antonym []string `xml:"ant"`

	// Information about the field of application of the entry/sense.
	// When absent, general application is implied. Entity coding for
	Field []Field `xml:"field"`

	// This element is used for other relevant information about
	// the entry/sense. As with part-of-speech, information will usually
	// apply to several senses.
	Misc []Misc `xml:"misc"`

	// The sense-information elements provided for additional
	// information to be recorded about a sense. Typical usage would
	// be to indicate such things as level of currency of a sense, the
	// regional variations, etc.
	Info []string `xml:"s_inf"`

	// This element records the information about the source
	// language(s) of a loan-word/gairaigo. If the source language is other
	// than English, the language is indicated by the xml:lang attribute.
	// The element value (if any) is the source word or phrase.
	LSource []LSource `xml:"lsource"`
	Dialect []Dialect `xml:"dial"`
	Gloss   []string  `xml:"gloss"`
	Example []string  `xml:"example"`
}

type LSource struct {
	XMLName xml.Name `xml:"lsource"`
	Lang    string   `xml:"xml:lang,attr"`
	Type    string   `xml:"ls_type,attr"`
	Wasei   string   `xml:"ls_wasei,attr"`
	Source  string   `xml:",chardata"`
}

var positionDescriptions map[Position]string = map[Position]string{
	AdjI:              "adjective (keiyoushi)",
	AdjKu:             "`ku' adjective (archaic)",
	AdjNa:             "adjectival nouns or quasi-adjectives (keiyodoshi)",
	AdjNari:           "archaic/formal form of na-adjective",
	AdjNo:             "nouns which may take the genitive case particle `no'",
	AdjPreNoun:        "pre-noun adjectival (rentaishi)",
	AdjShiku:          "`shiku' adjective (archaic)",
	AdjTaru:           "`taru' adjective",
	AdjFunctional:     "noun or verb acting prenominally",
	Adv:               "adverb (fukushi)",
	AdvTo:             "adverb taking the `to' particle",
	Aux:               "auxiliary",
	AuxAdj:            "auxiliary adjective",
	AuxVerb:           "auxiliary verb",
	Conj:              "conjunction",
	Counter:           "counter",
	Expression:        "Expressions (phrases, clauses, etc.)",
	Interjection:      "interjection (kandoushi)",
	Noun:              "noun (common) (futsuumeishi)",
	NounAdv:           "adverbial noun (fukushitekimeishi)",
	NounPropper:       "proper noun",
	NounPrefix:        "noun, used as a prefix",
	NounSuffix:        "noun, used as a suffix",
	NounTemporal:      "noun (temporal) (jisoumeishi)",
	Numeric:           "numeric",
	Pronoun:           "pronoun",
	Prefix:            "prefix",
	Particle:          "particle",
	Suffix:            "suffix",
	Verb1:             "Ichidan verb",
	Verb2as:           "Nidan verb with 'u' ending (archaic)",
	Verb2bk:           "Nidan verb (upper class) with `bu' ending (archaic)",
	Verb2ds:           "Nidan verb (lower class) with `dzu' ending (archaic)",
	Verb2gk:           "Nidan verb (upper class) with `gu' ending (archaic)",
	Verb2gs:           "Nidan verb (lower class) with `gu' ending (archaic)",
	Verb2hk:           "Nidan verb (upper class) with `hu/fu' ending (archaic)",
	Verb2hs:           "Nidan verb (lower class) with `hu/fu' ending (archaic)",
	Verb2kk:           "Nidan verb (upper class) with `ku' ending (archaic)",
	Verb2ks:           "Nidan verb (lower class) with `ku' ending (archaic)",
	Verb2ms:           "Nidan verb (lower class) with `mu' ending (archaic)",
	Verb2ns:           "Nidan verb (lower class) with `nu' ending (archaic)",
	Verb2rk:           "Nidan verb (upper class) with `ru' ending (archaic)",
	Verb2rs:           "Nidan verb (lower class) with `ru' ending (archaic)",
	Verb2ss:           "Nidan verb (lower class) with `su' ending (archaic)",
	Verb2tk:           "Nidan verb (upper class) with `tsu' ending (archaic)",
	Verb2ts:           "Nidan verb (lower class) with `tsu' ending (archaic)",
	Verb2ws:           "Nidan verb (lower class) with `u' ending and `we' conjugation (archaic)",
	Verb2yk:           "Nidan verb (upper class) with `yu' ending (archaic)",
	Verb2ys:           "Nidan verb (lower class) with `yu' ending (archaic)",
	Verb2zs:           "Nidan verb (lower class) with `zu' ending (archaic)",
	Verb4b:            "Yodan verb with `bu' ending (archaic)",
	Verb4h:            "Yodan verb with `hu/fu' ending (archaic)",
	Verb4k:            "Yodan verb with `ku' ending (archaic)",
	Verb4r:            "Yodan verb with `ru' ending (archaic)",
	Verb4s:            "Yodan verb with `su' ending (archaic)",
	Verb4t:            "Yodan verb with `tsu' ending (archaic)",
	Verb5aru:          "Godan verb - -aru special class",
	Verb5b:            "Godan verb with `bu' ending",
	Verb5g:            "Godan verb with `gu' ending",
	Verb5k:            "Godan verb with `ku' ending",
	Verb5ks:           "Godan verb - Iku/Yuku special class",
	Verb5m:            "Godan verb with `mu' ending",
	Verb5n:            "Godan verb with `nu' ending",
	Verb5r:            "Godan verb with `ru' ending",
	Verb5ri:           "Godan verb with `ru' ending (irregular verb)",
	Verb5s:            "Godan verb with `su' ending",
	Verb5t:            "Godan verb with `tsu' ending",
	Verb5u:            "Godan verb with `u' ending",
	Verb5us:           "Godan verb with `u' ending (special class)",
	VerbIntransitive:  "intransitive verb",
	VerbKuru:          "Kuru verb - special class",
	VerbNu:            "irregular nu verb",
	VerbRu:            "irregular ru verb, plain form ends with -ri",
	VerbSuru:          "noun or participle which takes the aux. verb suru",
	VerbSu:            "su verb - precursor to the modern suru",
	VerbIrregularSuru: "suru verb - irregular",
	VerbSuruSpecial:   "suru verb - special class",
	VerbTransitive:    "transitive verb",
	VerbZuru:          "Ichidan verb - zuru verb (alternative form of -jiru verbs)"}

var fieldDescriptions map[Field]string = map[Field]string{
	Anatomical:   "anatomical",
	Architecture: "architecture",
	Astronomy:    "astronomy",
	Baseball:     "baseball",
	Biology:      "biology",
	Botany:       "botany",
	buddhist:     "Buddhist",
	Business:     "business",
	Chemistry:    "chemistry",
	Computer:     "computers",
	Economics:    "economics",
	Engineering:  "engineering",
	Finance:      "finance",
	Food:         "food",
	Geology:      "geology",
	Geometry:     "geometry",
	Law:          "law",
	Linguistics:  "linguistics",
	Martial:      "martial arts",
	Mathematics:  "mathematics",
	Medicine:     "medicine",
	Military:     "military",
	Music:        "music",
	Physics:      "physics",
	shinto:       "Shinto",
	Sports:       "sports",
	Sumo:         "sumo",
	Zoology:      "zoology"}

var miscDescriptions map[Misc]string = map[Misc]string{
	Abbreviation:   "abbreviation",
	Archaism:       "archaism",
	ChildLanguage:  "children's language",
	Colloquialism:  "colloquialism",
	Derogatory:     "derogatory",
	Familiar:       "familiar language",
	FemaleLanguage: "female term or language",
	Honorific:      "honorific or respectful (sonkeigo) language",
	Humble:         "humble (kenjougo) language",
	Idiomatic:      "idiomatic expression",
	Jocular:        "jocular, humorous term",
	MaleLanguage:   "male term or language",
	Manga:          "manga slang",
	Obsolete:       "obsolete term",
	Obscure:        "obscure term",
	Onomatopoeic:   "onomatopoeic or mimetic word",
	Poetical:       "poetical term",
	Polite:         "polite (teineigo) language",
	Proverb:        "proverb",
	Rare:           "rare",
	Sensitive:      "sensitive",
	Slang:          "slang",
	KanaAlone:      "word usually written using kana alone",
	Vulgar:         "vulgar expression or word",
	XRated:         "rude or X-rated term (not displayed in educational software)"}

var orthographyDescriptions map[Orthography]string = map[Orthography]string{
	Ateji:                   "ateji (phonetic) reading",
	Gikun:                   "gikun (meaning as reading) or jukujikun (special kanji reading)",
	IrregularKanji:          "word containing irregular kanji usage",
	IrregularKana:           "word containing irregular kana usage",
	IrregularOkurigana:      "irregular okurigana usage",
	OutdatedKanji:           "out-dated or obsolete kanji usage",
	OutdatedKana:            "out-dated or obsolete kana usage",
	OutdatedOrIrregularKana: "old or irregular kana form",
	KanjiAlone:              "word usually written using kanji alone"}

var dialectDescriptions map[Dialect]string = map[Dialect]string{
	HokkaidoBen: "Hokkaido-ben",
	KansaiBen:   "Kansai-ben",
	KantouBen:   "Kantou-ben",
	KyotoBen:    "Kyoto-ben",
	KyuushuuBen: "Kyuushuu-ben",
	NaganoBen:   "Nagano-ben",
	OsakaBen:    "Osaka-ben",
	RyuukyuuBen: "Ryuukyuu-ben",
	TouhokuBen:  "Touhoku-ben",
	TosaBen:     "Tosa-ben",
	TsugaruBen:  "Tsugaru-ben"}

// Produce a human readable description of a Position
func DescribePosition(position Position) string {
	return positionDescriptions[position]
}

// Produce a human readable description of a Field
func DescribeField(field Field) string {
	return fieldDescriptions[field]
}

// Produce a human readable description of an Orthography
func DescribeOrthography(orthography Orthography) string {
	return orthographyDescriptions[orthography]
}

// Produce a human readable description of a Dialect
func DescribeDialect(dialect Dialect) string {
	return dialectDescriptions[dialect]
}

// Produce a human readable description of a Misc entity
func describeMisc(misc Misc) string {
	return miscDescriptions[misc]
}

func Read(r io.Reader) (JMDict, error) {
	var dict JMDict
	decoder := xml.NewDecoder(r)

	decoder.Entity = make(map[string]string)
	for k := range positionDescriptions {
		decoder.Entity[string(k)] = string(k)
	}
	for k := range fieldDescriptions {
		decoder.Entity[string(k)] = string(k)
	}
	for k := range miscDescriptions {
		decoder.Entity[string(k)] = string(k)
	}
	for k := range orthographyDescriptions {
		decoder.Entity[string(k)] = string(k)
	}
	for k := range dialectDescriptions {
		decoder.Entity[string(k)] = string(k)
	}

	err := decoder.Decode(&dict)
	return dict, err
}
