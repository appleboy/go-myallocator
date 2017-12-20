package ma

// Channel struct
type Channel struct {
	Status   string `json:"status"`
	Currency string `json:"currency"`
}

// PropertyChannelListRes for PropertyChannelList response
// missing some field
type PropertyChannelListRes struct {
	Properties []struct {
		PropertyID int `json:"PropertyId"`
		Channels   struct {
			Koh    Channel `json:"koh"`
			Asyo   Channel `json:"asyo"`
			Vjo    Channel `json:"vjo"`
			Bb     Channel `json:"bb"`
			Mal    Channel `json:"mal"`
			Outp   Channel `json:"outp"`
			Trav   Channel `json:"trav"`
			Mcs    Channel `json:"mcs"`
			Ago    Channel `json:"ago"`
			Hcu2   Channel `json:"hcu2"`
			Buen   Channel `json:"buen"`
			Myb    Channel `json:"myb"`
			Bkb    Channel `json:"bkb"`
			Ptch   Channel `json:"ptch"`
			Bp     Channel `json:"bp"`
			Lmin   Channel `json:"lmin"`
			G2I    Channel `json:"g2i"`
			Htrv   Channel `json:"htrv"`
			Clt    Channel `json:"clt"`
			Han    Channel `json:"han"`
			Yoho   Channel `json:"yoho"`
			Hde    Channel `json:"hde"`
			Wanu   Channel `json:"wanu"`
			Hd     Channel `json:"hd"`
			Orb    Channel `json:"orb"`
			Rrep   Channel `json:"rrep"`
			Loc    Channel `json:"loc"`
			Air    Channel `json:"air"`
			Deco   Channel `json:"deco"`
			Btob   Channel `json:"btob"`
			Enz    Channel `json:"enz"`
			Tpvl   Channel `json:"tpvl"`
			Hbe    Channel `json:"hbe"`
			Saas   Channel `json:"saas"`
			Liko   Channel `json:"liko"`
			Fk     Channel `json:"fk"`
			Hhb    Channel `json:"hhb"`
			Buun   Channel `json:"buun"`
			Ost    Channel `json:"ost"`
			Hc     Channel `json:"hc"`
			Gom    Channel `json:"gom"`
			Vhh    Channel `json:"vhh"`
			Prsc   Channel `json:"prsc"`
			Eb     Channel `json:"eb"`
			NineFl Channel `json:"9fl"`
			Mrbb   Channel `json:"mrbb"`
			Pool   Channel `json:"pool"`
			Bole   Channel `json:"bole"`
			Whoo   Channel `json:"whoo"`
			Eget   Channel `json:"eget"`
			Hotu   Channel `json:"hotu"`
			Inst   Channel `json:"inst"`
			Hs     Channel `json:"hs"`
			Grup   Channel `json:"grup"`
			Xmlt   Channel `json:"xmlt"`
			Vaca   Channel `json:"vaca"`
			Net    Channel `json:"net"`
			Tmd    Channel `json:"tmd"`
			Loop   Channel `json:"loop"`
			Bon    Channel `json:"bon"`
			Hoga   Channel `json:"hoga"`
			Ct     Channel `json:"ct"`
			Kch    Channel `json:"kch"`
			Art    Channel `json:"art"`
			Tlok   Channel `json:"tlok"`
			Gr     Channel `json:"gr"`
			Trpz   Channel `json:"trpz"`
			Hb2    Channel `json:"hb2"`
			Eba    Channel `json:"eba"`
			Hort   Channel `json:"hort"`
			Rep    Channel `json:"rep"`
			Bnw    Channel `json:"bnw"`
			Vcay   Channel `json:"vcay"`
			Cwd    Channel `json:"cwd"`
			Hrs    Channel `json:"hrs"`
			Tri    Channel `json:"tri"`
			Apal   Channel `json:"apal"`
			Boo    Channel `json:"boo"`
			Snrg   Channel `json:"snrg"`
			Lako   Channel `json:"lako"`
			Toqt   Channel `json:"toqt"`
			Chot   Channel `json:"chot"`
			Pre    Channel `json:"pre"`
			Cozy   Channel `json:"cozy"`
			Advt   Channel `json:"advt"`
			Wim    Channel `json:"wim"`
			Adve   Channel `json:"adve"`
			Tob    Channel `json:"tob"`
			Ho     Channel `json:"ho"`
			Hopa   Channel `json:"hopa"`
			Flh    Channel `json:"flh"`
			Hipc   Channel `json:"hipc"`
			Hoca   Channel `json:"hoca"`
			Rtel   Channel `json:"rtel"`
			Mmt    Channel `json:"mmt"`
			Enbc   Channel `json:"enbc"`
			Lr     Channel `json:"lr"`
			Bnb    Channel `json:"bnb"`
			Rec    Channel `json:"rec"`
			H2H    Channel `json:"h2h"`
			Hand   Channel `json:"hand"`
			Hmi    Channel `json:"hmi"`
			Bnb2   Channel `json:"bnb2"`
			Sbkr   Channel `json:"sbkr"`
			Rfc    Channel `json:"rfc"`
			Pand   Channel `json:"pand"`
			Kqbe   Channel `json:"kqbe"`
			Bud    Channel `json:"bud"`
			Ini    Channel `json:"ini"`
			Adv    Channel `json:"adv"`
			Cs     Channel `json:"cs"`
			Ysh    Channel `json:"ysh"`
			Etb    Channel `json:"etb"`
			Exp    Channel `json:"exp"`
			Evry   Channel `json:"evry"`
			Wix    Channel `json:"wix"`
			Luna   Channel `json:"luna"`
			Rsrv   Channel `json:"rsrv"`
			Vrbo   Channel `json:"vrbo"`
			Bd     Channel `json:"bd"`
			Bh     Channel `json:"bh"`
			All    Channel `json:"all"`
			R2N    Channel `json:"r2n"`
			Hw     Channel `json:"hw"`
			Bal    Channel `json:"bal"`
			Bv     Channel `json:"bv"`
			Fam    Channel `json:"fam"`
			Hr     Channel `json:"hr"`
			Ote    Channel `json:"ote"`
			Tavr   Channel `json:"tavr"`
			Mega   Channel `json:"mega"`
			Bdrm   Channel `json:"bdrm"`
			Odi    Channel `json:"odi"`
			Iwb    Channel `json:"iwb"`
			Ddhm   Channel `json:"ddhm"`
			Ht     Channel `json:"ht"`
			Skd    Channel `json:"skd"`
			Hcu    Channel `json:"hcu"`
			Hola   Channel `json:"hola"`
			Sawa   Channel `json:"sawa"`
			Tom    Channel `json:"tom"`
			Esc    Channel `json:"esc"`
			Cb     Channel `json:"cb"`
			Ohma   Channel `json:"ohma"`
			Ya     Channel `json:"ya"`
			Sole   Channel `json:"sole"`
			Aveo   Channel `json:"aveo"`
			Coco   Channel `json:"coco"`
			Iper   Channel `json:"iper"`
			Bbit   Channel `json:"bbit"`
			Adsh   Channel `json:"adsh"`
			Ktel   Channel `json:"ktel"`
			Akt    Channel `json:"akt"`
			Lvur   Channel `json:"lvur"`
			Nom    Channel `json:"nom"`
			Hi     Channel `json:"hi"`
			Vive   Channel `json:"vive"`
			Lsbv   Channel `json:"lsbv"`
			Stya   Channel `json:"stya"`
			Wpk    Channel `json:"wpk"`
			Sur    Channel `json:"sur"`
			Zldv   Channel `json:"zldv"`
			Kni    Channel `json:"kni"`
			Gta    Channel `json:"gta"`
			Max    Channel `json:"max"`
			Goi    Channel `json:"goi"`
			Hb     Channel `json:"hb"`
			Ctr    Channel `json:"ctr"`
			Hom    Channel `json:"hom"`
			Bbnl   Channel `json:"bbnl"`
			Lota   Channel `json:"lota"`
			Lmg    Channel `json:"lmg"`
			Ta     Channel `json:"ta"`
			Bool   Channel `json:"bool"`
			Trab   Channel `json:"trab"`
			Ibc    Channel `json:"ibc"`
			Oa     Channel `json:"oa"`
			Hotq   Channel `json:"hotq"`
			Pzg    Channel `json:"pzg"`
			Hou    Channel `json:"hou"`
			Hw2    Channel `json:"hw2"`
		} `json:"Channels"`
	} `json:"Properties"`
	Success bool `json:"Success"`
}
