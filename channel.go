package ma

// PropertyChannelListRes for PropertyChannelList response
// missing some field
type PropertyChannelListRes struct {
	Properties []struct {
		PropertyID int `json:"PropertyId"`
		Channels   struct {
			Koh struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"koh"`
			Asyo struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"asyo"`
			Vjo struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"vjo"`
			Bb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bb"`
			Mal struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"mal"`
			Outp struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"outp"`
			Trav struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"trav"`
			Mcs struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"mcs"`
			Ago struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ago"`
			Hcu2 struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hcu2"`
			Buen struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"buen"`
			Myb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"myb"`
			Bkb struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bkb"`
			Ptch struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ptch"`
			Bp struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bp"`
			Lmin struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"lmin"`
			G2I struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"g2i"`
			Htrv struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"htrv"`
			Clt struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"clt"`
			Han struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"han"`
			Yoho struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"yoho"`
			Hde struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hde"`
			Wanu struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"wanu"`
			Hd struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hd"`
			Orb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"orb"`
			Rrep struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"rrep"`
			Loc struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"loc"`
			Air struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"air"`
			Deco struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"deco"`
			Btob struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"btob"`
			Enz struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"enz"`
			Tpvl struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"tpvl"`
			Hbe struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hbe"`
			Saas struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"saas"`
			Liko struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"liko"`
			Fk struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"fk"`
			Hhb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hhb"`
			Buun struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"buun"`
			Ost struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ost"`
			Hc struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hc"`
			Gom struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"gom"`
			Vhh struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"vhh"`
			Prsc struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"prsc"`
			Eb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"eb"`
			NineFl struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"9fl"`
			Mrbb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"mrbb"`
			Pool struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"pool"`
			Bole struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bole"`
			Whoo struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"whoo"`
			Eget struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"eget"`
			Hotu struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hotu"`
			Inst struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"inst"`
			Hs struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hs"`
			Grup struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"grup"`
			Xmlt struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"xmlt"`
			Vaca struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"vaca"`
			Net struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"net"`
			Tmd struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"tmd"`
			Loop struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"loop"`
			Bon struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bon"`
			Hoga struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hoga"`
			Ct struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ct"`
			Kch struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"kch"`
			Art struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"art"`
			Tlok struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"tlok"`
			Gr struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"gr"`
			Trpz struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"trpz"`
			Hb2 struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hb2"`
			Eba struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"eba"`
			Hort struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hort"`
			Rep struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"rep"`
			Bnw struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bnw"`
			Vcay struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"vcay"`
			Cwd struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"cwd"`
			Hrs struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hrs"`
			Tri struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"tri"`
			Apal struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"apal"`
			Boo struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"boo"`
			Snrg struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"snrg"`
			Lako struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"lako"`
			Toqt struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"toqt"`
			Chot struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"chot"`
			Pre struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"pre"`
			Cozy struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"cozy"`
			Advt struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"advt"`
			Wim struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"wim"`
			Adve struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"adve"`
			Tob struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"tob"`
			Ho struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ho"`
			Hopa struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hopa"`
			Flh struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"flh"`
			Hipc struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hipc"`
			Hoca struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hoca"`
			Rtel struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"rtel"`
			Mmt struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"mmt"`
			Enbc struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"enbc"`
			Lr struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"lr"`
			Bnb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bnb"`
			Rec struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"rec"`
			H2H struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"h2h"`
			Hand struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hand"`
			Hmi struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hmi"`
			Bnb2 struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bnb2"`
			Sbkr struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"sbkr"`
			Rfc struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"rfc"`
			Pand struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"pand"`
			Kqbe struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"kqbe"`
			Bud struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bud"`
			Ini struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ini"`
			Adv struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"adv"`
			Cs struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"cs"`
			Ysh struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ysh"`
			Etb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"etb"`
			Exp struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"exp"`
			Evry struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"evry"`
			Wix struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"wix"`
			Luna struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"luna"`
			Rsrv struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"rsrv"`
			Vrbo struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"vrbo"`
			Bd struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bd"`
			Bh struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bh"`
			All struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"all"`
			R2N struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"r2n"`
			Hw struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hw"`
			Bal struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bal"`
			Bv struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bv"`
			Fam struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"fam"`
			Hr struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hr"`
			Ote struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ote"`
			Tavr struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"tavr"`
			Mega struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"mega"`
			Bdrm struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bdrm"`
			Odi struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"odi"`
			Iwb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"iwb"`
			Ddhm struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ddhm"`
			Ht struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ht"`
			Skd struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"skd"`
			Hcu struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hcu"`
			Hola struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hola"`
			Sawa struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"sawa"`
			Tom struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"tom"`
			Esc struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"esc"`
			Cb struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"cb"`
			Ohma struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ohma"`
			Ya struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"ya"`
			Sole struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"sole"`
			Aveo struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"aveo"`
			Coco struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"coco"`
			Iper struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"iper"`
			Bbit struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"bbit"`
			Adsh struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"adsh"`
			Ktel struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ktel"`
			Akt struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"akt"`
			Lvur struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"lvur"`
			Nom struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"nom"`
			Hi struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hi"`
			Vive struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"vive"`
			Lsbv struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"lsbv"`
			Stya struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"stya"`
			Wpk struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"wpk"`
			Sur struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"sur"`
			Zldv struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"zldv"`
			Kni struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"kni"`
			Gta struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"gta"`
			Max struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"max"`
			Goi struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"goi"`
			Hb struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hb"`
			Ctr struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ctr"`
			Hom struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hom"`
			Bbnl struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bbnl"`
			Lota struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"lota"`
			Lmg struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"lmg"`
			Ta struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ta"`
			Bool struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"bool"`
			Trab struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"trab"`
			Ibc struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"ibc"`
			Oa struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"oa"`
			Hotq struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hotq"`
			Pzg struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"pzg"`
			Hou struct {
				Status   string `json:"status"`
				Currency string `json:"currency"`
			} `json:"hou"`
			Hw2 struct {
				Currency string `json:"currency"`
				Status   string `json:"status"`
			} `json:"hw2"`
		} `json:"Channels"`
	} `json:"Properties"`
	Success bool `json:"Success"`
}
