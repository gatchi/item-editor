package main

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

