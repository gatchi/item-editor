package itempmt

type PMTArmor struct {
	id            uint32
	model          int16
	skin           int16
	teampoints     int32
	dfp            int16
	evp            int16
	blockparticle uint8
	blockeffect   uint8
	_class        uint8
    reserved1     uint8 
	requiredlevel uint8
	efr           uint8
	eth           uint8
	eic           uint8
	edk           uint8
	elt           uint8
	dfprange      uint8
	evprange      uint8
	statboost     uint8
	techboost     uint8
	unknown1       int16
}

type PMTShield PMTArmor

type PMTUnit struct {
	id            uint32
	model          int16
	skin           int16
	teampoints     int32
    stat           int16
	statamount     int16
    plusminus      int8
    reserved    [3]int8
}

type Mag struct {
	id            uint32
	model          int16
	skin           int16
	teampoints     int32
	feedtable     uint8
	unknown1      uint8
	photonblast   uint8
	activation    uint8
	pbfull        uint8
	lowhp         uint8
	death         uint8
	boss          uint8
	pbfullflag    uint8
	lowhpflag     uint8
	deathflag     uint8
	bossflag      uint8
    _class        uint8
	reserved   [3]uint8
}

type Item struct {
	id            uint32
	model          int16
	skin           int16
	teampoints     int32
	amount         int16
	tech           int16
	cost           int32
	itemflag      uint8
	reserved   [3]uint8
}

type PMTWeapon struct {
	id            uint32
	model          int16
	skin           int16
	teampoints     int32
	_class        uint8
	reserved      uint8
	atpmin         int16
	atpmax         int16
	atpreq         int16
	mstreq         int16
	atareq         int16
	mst            int16
	maxgrind      uint8
	photon         int8
	special       uint8
	ata           uint8
	statboost     uint8
	projectile    uint8
	photontrail1x  int8
	photontrail1y  int8
	photontrail2x  int8
	photontrail2y  int8
	photontype     int8
	unknown1       int8
	unknown2       int8
	unknown3       int8
	unknown4       int8
	techboost     uint8
	combotype     uint8
}

type MagFeedItem struct {
	defence      int8
	power        int8
	dexterity    int8
	mind         int8
	iq           int8
	sync         int8
	reserved  [3]int8
}

type MagFeed struct {
	item  [11]MagFeedItem
}

type MagFeedTable struct {
	table [8]MagFeed
}

type StarValue struct {
	starvalue uint8
}

type Special struct {
	Type   int16
	amount int16
}

type StatBoost struct {
	stat1 uint8
	stat2 uint8
	amount1 int16
	amount2 int16
}

type TechUseClass struct {
	_class [12]uint8
}

type TechUseTable struct {
	tech [19]TechUseClass
}

type Combination struct {
	useitem       [3]uint8
	equippeditem  [3]uint8
	itemresult    [3]uint8
	maglevel         uint8
	grind            uint8
	level            uint8
	_class           uint8
	reserved      [3]uint8
}

type TechBoost struct {
	tech1   int32
	boost1  float32
	tech2   int32
	boost2	float32
	tech3   int32
	boost3  float32
}

type EventItem struct {
	item   [3]uint8
	chance    uint8
}

type Unsealable struct {
	item   [3]uint8
	reserved  uint8
}

type WeaponSaleDivisor struct {
	saledivisor  [165]float32
}

type SaleDivisor struct {
	armor  float32
	shield float32
	unit   float32
	mag    float32
}

type AttackAnimation struct {
	animation uint8
}

type LookupTableItem struct {
	pointer int32
}

type ItemCount struct {
	weaponcount        uint32
	armorcount         uint32
	shieldcount        uint32
	unitcount          uint32
	itemcount          uint32
	magcount           uint32
	starvaluecount     uint32
	specialcount       uint32
	statboostcount     uint32
	combinationcount   uint32
	techboostcount     uint32
	unwrapcount        uint32
	xmasitemcount      uint32
	easteritemcount    uint32
	halloweenitemcount uint32
	unsealablecount    uint32
}

type ItemPMT struct {
	itemcount           ItemCount
	pmtlookuptbl        [0x04][0xFF][0xFF]LookupTableItem
	magfeedtbl         *MagFeedTable
	starvaluetable     *StarValue
	starvalueoffset     uint32
	specialtbl         *Special
	statboosttbl       *StatBoost
	techusetbl         *TechUseTable
	combinationtbl     *Combination
	techboosttbl       *TechBoost
	attackanimationtbl *AttackAnimation
	xmastitemtbl       *EventItem
	easteritemtbl      *EventItem
	halloweenitemtbl   *EventItem
	unsealabletbl      *Unsealable
	weaponsaledivtbl   *WeaponSaleDivisor
	saledivtbl         *SaleDivisor
}
