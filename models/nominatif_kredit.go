package models

import (
	"database/sql"
	"strings"
)

// FilterParams represents filter parameters
type FilterParams struct {
	CAB            string `json:"cab"`
	AO             string `json:"ao"`
	KET_KD_PRD     string `json:"ket_kd_prd"`
	TEMPAT_BEKERJA string `json:"tempat_bekerja"`
}

type NominatifKredit struct {
	ID                   int             `json:"id" db:"id"`
	DATADATE             *string         `json:"DATADATE" db:"DATADATE"`
	CAB                  *string         `json:"CAB" db:"CAB"`
	NOMOR_REKENING       *string         `json:"NOMOR_REKENING" db:"NOMOR_REKENING"`
	NO_CIF               *string         `json:"NO_CIF" db:"NO_CIF"`
	NAMA_NASABAH         *string         `json:"NAMA_NASABAH" db:"NAMA_NASABAH"`
	ALAMAT               *string         `json:"ALAMAT" db:"ALAMAT"`
	KODE_KOLEK           *string         `json:"KODE_KOLEK" db:"KODE_KOLEK"`
	JML_HRI_PKK          *int            `json:"JML_HRI_PKK" db:"JML_HRI_PKK"`
	JML_HARI_BGA         *int            `json:"JML_HARI_BGA" db:"JML_HARI_BGA"`
	JML_HARI_TUNGGAKAN   *int            `json:"JML_HARI_TUNGGAKAN" db:"JML_HARI_TUNGGAKAN"`
	KD_PRD               *string         `json:"KD_PRD" db:"KD_PRD"`
	KET_KD_PRD           *string         `json:"KET_KD_PRD" db:"KET_KD_PRD"`
	NOMOR_PERJANJIAN     *string         `json:"NOMOR_PERJANJIAN" db:"NOMOR_PERJANJIAN"`
	NO_AKSEP             *string         `json:"NO_AKSEP" db:"NO_AKSEP"`
	TGL_PK               *string         `json:"TGL_PK" db:"TGL_PK"`
	TGL_AWAL_FAS         *string         `json:"TGL_AWAL_FAS" db:"TGL_AWAL_FAS"`
	TGL_AKHIR_FAS        *string         `json:"TGL_AKHIR_FAS" db:"TGL_AKHIR_FAS"`
	TGL_AWAL_AKSEP       *string         `json:"TGL_AWAL_AKSEP" db:"TGL_AWAL_AKSEP"`
	TGL_AKH_AKSEP        *string         `json:"TGL_AKH_AKSEP" db:"TGL_AKH_AKSEP"`
	PLAFOND_AWAL         *float64        `json:"PLAFOND_AWAL" db:"PLAFOND_AWAL"`
	BAKI_DEBET           *float64        `json:"BAKI_DEBET" db:"BAKI_DEBET"`
	LONGGAR_TARIK        *float64        `json:"LONGGAR_TARIK" db:"LONGGAR_TARIK"`
	BGA                  *float64        `json:"BGA" db:"BGA"`
	TUNGGAKAN_POKOK      *float64        `json:"TUNGGAKAN_POKOK" db:"TUNGGAKAN_POKOK"`
	TUNGGAKAN_BUNGA      *float64        `json:"TUNGGAKAN_BUNGA" db:"TUNGGAKAN_BUNGA"`
	BGA_JTH_TEMPO        *float64        `json:"BGA_JTH_TEMPO" db:"BGA_JTH_TEMPO"`
	SMP_TGL_CADANG       *string         `json:"SMP_TGL_CADANG" db:"SMP_TGL_CADANG"`
	NILAI_CADANG         *float64        `json:"NILAI_CADANG" db:"NILAI_CADANG"`
	ANGSURAN_TOTAL       *float64        `json:"ANGSURAN_TOTAL" db:"ANGSURAN_TOTAL"`
	TGL_PROSES_DENDA     *string         `json:"TGL_PROSES_DENDA" db:"TGL_PROSES_DENDA"`
	AKUM_DENDA_PKK       *float64        `json:"AKUM_DENDA_PKK" db:"AKUM_DENDA_PKK"`
	AKUM_DENDA_BGA       *float64        `json:"AKUM_DENDA_BGA" db:"AKUM_DENDA_BGA"`
	PRD_AMORT            *float64        `json:"PRD_AMORT" db:"PRD_AMORT"`
	PRDK_AMORT           *float64        `json:"PRDK_AMORT" db:"PRDK_AMORT"`
	FLAG                 *string         `json:"FLAG" db:"FLAG"`
	TGL_AMORT            *string         `json:"TGL_AMORT" db:"TGL_AMORT"`
	NILAI_BIAYA_PROVISI  *float64        `json:"NILAI_BIAYA_PROVISI" db:"NILAI_BIAYA_PROVISI"`
	AMORTISASI_PER_PRD   *float64        `json:"AMORTISASI_PER_PRD" db:"AMORTISASI_PER_PRD"`
	SISA_AMORT_PROV      *float64        `json:"SISA_AMORT_PROV" db:"SISA_AMORT_PROV"`
	TAGIH_BIAYA_PROV     *float64        `json:"TAGIH_BIAYA_PROV" db:"TAGIH_BIAYA_PROV"`
	NILAI_BIAYA_ADM      *float64        `json:"NILAI_BIAYA_ADM" db:"NILAI_BIAYA_ADM"`
	AMORT_ADM_PER_PRD    *float64        `json:"AMORT_ADM_PER_PRD" db:"AMORT_ADM_PER_PRD"`
	SISA_AMORT_ADM       *float64        `json:"SISA_AMORT_ADM" db:"SISA_AMORT_ADM"`
	BYA_ASURANSI         *float64        `json:"BYA_ASURANSI" db:"BYA_ASURANSI"`
	BYA_NOTARIS          *float64        `json:"BYA_NOTARIS" db:"BYA_NOTARIS"`
	PKK_JATEM            *float64        `json:"PKK_JATEM" db:"PKK_JATEM"`
	BGA_JATEM            *float64        `json:"BGA_JATEM" db:"BGA_JATEM"`
	REK_BYR_PKK_BGA      *string         `json:"REK_BYR_PKK_BGA" db:"REK_BYR_PKK_BGA"`
	SLD_REK_DB           *string         `json:"SLD_REK_DB" db:"SLD_REK_DB"`
	KD_INSTANSI          *string         `json:"KD_INSTANSI" db:"KD_INSTANSI"`
	NM_INSTANSI          *string         `json:"NM_INSTANSI" db:"NM_INSTANSI"`
	REK_BENDAHARA        *string         `json:"REK_BENDAHARA" db:"REK_BENDAHARA"`
	SFT_KRD              *string         `json:"SFT_KRD" db:"SFT_KRD"`
	GOL_KRD              *string         `json:"GOL_KRD" db:"GOL_KRD"`
	JNS_KRD              *string         `json:"JNS_KRD" db:"JNS_KRD"`
	SKTR_EKNM            *string         `json:"SKTR_EKNM" db:"SKTR_EKNM"`
	ORNTS                *string         `json:"ORNTS" db:"ORNTS"`
	NO_HP                *string         `json:"NO_HP" db:"NO_HP"`
	POKOK_PINJAMAN       *float64        `json:"POKOK_PINJAMAN" db:"POKOK_PINJAMAN"`
	TITIPAN_EFEKTIF      *string         `json:"TITIPAN_EFEKTIF" db:"TITIPAN_EFEKTIF"`
	JANGKA_WAKTU         *int            `json:"JANGKA_WAKTU" db:"JANGKA_WAKTU"`
	REK_PENCAIRAN        *string         `json:"REK_PENCAIRAN" db:"REK_PENCAIRAN"`
	NO_REKENING_LAMA     *string         `json:"NO_REKENING_LAMA" db:"NO_REKENING_LAMA"`
	CIF_LAMA             *string         `json:"CIF_LAMA" db:"CIF_LAMA"`
	KODE_GROUP           *string         `json:"KODE_GROUP" db:"KODE_GROUP"`
	KET_GROUP            *string         `json:"KET_GROUP" db:"KET_GROUP"`
	TGL_LAHIR            *string         `json:"TGL_LAHIR" db:"TGL_LAHIR"`
	NIK                  *string         `json:"NIK" db:"NIK"`
	NIP                  *string         `json:"NIP" db:"NIP"`
	NILAI_BYA_TRANS      *float64        `json:"NILAI_BYA_TRANS" db:"NILAI_BYA_TRANS"`
	AMORT_TRANS_PER_PRD  *float64        `json:"AMORT_TRANS_PER_PRD" db:"AMORT_TRANS_PER_PRD"`
	SISA_AMORT_TRANS     *float64        `json:"SISA_AMORT_TRANS" db:"SISA_AMORT_TRANS"`
	AO                   *string         `json:"AO" db:"AO"`
	CAB_REK              *string         `json:"CAB_REK" db:"CAB_REK"`
	KELURAHAN            *string         `json:"KELURAHAN" db:"KELURAHAN"`
	KECAMATAN            *string         `json:"KECAMATAN" db:"KECAMATAN"`
	CADANGAN_PPAP        *float64        `json:"CADANGAN_PPAP" db:"CADANGAN_PPAP"`
	TEMPAT_BEKERJA       *string         `json:"TEMPAT_BEKERJA" db:"TEMPAT_BEKERJA"`
	TGL_AKHIR_BAYAR      *string         `json:"TGL_AKHIR_BAYAR" db:"TGL_AKHIR_BAYAR"`
	PIHAK_TERKAIT        *string         `json:"PIHAK_TERKAIT" db:"PIHAK_TERKAIT"`
	JENIS_JAMINAN        *string         `json:"JENIS_JAMINAN" db:"JENIS_JAMINAN"`
	NILAI_LEGALITAS      *float64        `json:"NILAI_LEGALITAS" db:"NILAI_LEGALITAS"`
	RESTRUKTUR_KE        *int            `json:"RESTRUKTUR_KE" db:"RESTRUKTUR_KE"`
	TGL_VALID_KOLEK      *string         `json:"TGL_VALID_KOLEK" db:"TGL_VALID_KOLEK"`
	TGL_MACET            *string         `json:"TGL_MACET" db:"TGL_MACET"`
}

// NominatifKreditRepository handles database operations for nominatif_kredit
type NominatifKreditRepository struct {
	DB *sql.DB
}

// NewNominatifKreditRepository creates a new repository instance
func NewNominatifKreditRepository(db *sql.DB) *NominatifKreditRepository {
	return &NominatifKreditRepository{DB: db}
}

// GetAll retrieves all nominatif_kredit records with pagination
func (r *NominatifKreditRepository) GetAll(limit, offset int) ([]NominatifKredit, error) {
	query := `
		SELECT id, DATADATE, CAB, NOMOR_REKENING, NO_CIF, NAMA_NASABAH, ALAMAT, KODE_KOLEK,
		       JML_HRI_PKK, JML_HARI_BGA, JML_HARI_TUNGGAKAN, KD_PRD, KET_KD_PRD, NOMOR_PERJANJIAN,
		       NO_AKSEP, TGL_PK, TGL_AWAL_FAS, TGL_AKHIR_FAS, TGL_AWAL_AKSEP, TGL_AKH_AKSEP,
		       PLAFOND_AWAL, BAKI_DEBET, LONGGAR_TARIK, BGA, TUNGGAKAN_POKOK, TUNGGAKAN_BUNGA,
		       BGA_JTH_TEMPO, SMP_TGL_CADANG, NILAI_CADANG, ANGSURAN_TOTAL, TGL_PROSES_DENDA,
		       AKUM_DENDA_PKK, AKUM_DENDA_BGA, PRD_AMORT, PRDK_AMORT, FLAG, TGL_AMORT,
		       NILAI_BIAYA_PROVISI, AMORTISASI_PER_PRD, SISA_AMORT_PROV, TAGIH_BIAYA_PROV,
		       NILAI_BIAYA_ADM, AMORT_ADM_PER_PRD, SISA_AMORT_ADM, BYA_ASURANSI, BYA_NOTARIS,
		       PKK_JATEM, BGA_JATEM, REK_BYR_PKK_BGA, SLD_REK_DB, KD_INSTANSI, NM_INSTANSI,
		       REK_BENDAHARA, SFT_KRD, GOL_KRD, JNS_KRD, SKTR_EKNM, ORNTS, NO_HP,
		       POKOK_PINJAMAN, TITIPAN_EFEKTIF, JANGKA_WAKTU, REK_PENCAIRAN, NO_REKENING_LAMA,
		       CIF_LAMA, KODE_GROUP, KET_GROUP, TGL_LAHIR, NIK, NIP, NILAI_BYA_TRANS,
		       AMORT_TRANS_PER_PRD, SISA_AMORT_TRANS, AO, CAB_REK, KELURAHAN, KECAMATAN,
		       CADANGAN_PPAP, TEMPAT_BEKERJA, TGL_AKHIR_BAYAR, PIHAK_TERKAIT, JENIS_JAMINAN,
		       NILAI_LEGALITAS, RESTRUKTUR_KE, TGL_VALID_KOLEK, TGL_MACET
		FROM nominatif_kredit 
		ORDER BY id DESC 
		LIMIT ? OFFSET ?`

	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kredits []NominatifKredit
	for rows.Next() {
		var kredit NominatifKredit
		err := rows.Scan(
			&kredit.ID, &kredit.DATADATE, &kredit.CAB, &kredit.NOMOR_REKENING, &kredit.NO_CIF,
			&kredit.NAMA_NASABAH, &kredit.ALAMAT, &kredit.KODE_KOLEK, &kredit.JML_HRI_PKK,
			&kredit.JML_HARI_BGA, &kredit.JML_HARI_TUNGGAKAN, &kredit.KD_PRD, &kredit.KET_KD_PRD,
			&kredit.NOMOR_PERJANJIAN, &kredit.NO_AKSEP, &kredit.TGL_PK, &kredit.TGL_AWAL_FAS,
			&kredit.TGL_AKHIR_FAS, &kredit.TGL_AWAL_AKSEP, &kredit.TGL_AKH_AKSEP, &kredit.PLAFOND_AWAL,
			&kredit.BAKI_DEBET, &kredit.LONGGAR_TARIK, &kredit.BGA, &kredit.TUNGGAKAN_POKOK,
			&kredit.TUNGGAKAN_BUNGA, &kredit.BGA_JTH_TEMPO, &kredit.SMP_TGL_CADANG, &kredit.NILAI_CADANG,
			&kredit.ANGSURAN_TOTAL, &kredit.TGL_PROSES_DENDA, &kredit.AKUM_DENDA_PKK, &kredit.AKUM_DENDA_BGA,
			&kredit.PRD_AMORT, &kredit.PRDK_AMORT, &kredit.FLAG, &kredit.TGL_AMORT, &kredit.NILAI_BIAYA_PROVISI,
			&kredit.AMORTISASI_PER_PRD, &kredit.SISA_AMORT_PROV, &kredit.TAGIH_BIAYA_PROV, &kredit.NILAI_BIAYA_ADM,
			&kredit.AMORT_ADM_PER_PRD, &kredit.SISA_AMORT_ADM, &kredit.BYA_ASURANSI, &kredit.BYA_NOTARIS,
			&kredit.PKK_JATEM, &kredit.BGA_JATEM, &kredit.REK_BYR_PKK_BGA, &kredit.SLD_REK_DB,
			&kredit.KD_INSTANSI, &kredit.NM_INSTANSI, &kredit.REK_BENDAHARA, &kredit.SFT_KRD,
			&kredit.GOL_KRD, &kredit.JNS_KRD, &kredit.SKTR_EKNM, &kredit.ORNTS, &kredit.NO_HP,
			&kredit.POKOK_PINJAMAN, &kredit.TITIPAN_EFEKTIF, &kredit.JANGKA_WAKTU, &kredit.REK_PENCAIRAN,
			&kredit.NO_REKENING_LAMA, &kredit.CIF_LAMA, &kredit.KODE_GROUP, &kredit.KET_GROUP,
			&kredit.TGL_LAHIR, &kredit.NIK, &kredit.NIP, &kredit.NILAI_BYA_TRANS, &kredit.AMORT_TRANS_PER_PRD,
			&kredit.SISA_AMORT_TRANS, &kredit.AO, &kredit.CAB_REK, &kredit.KELURAHAN, &kredit.KECAMATAN,
			&kredit.CADANGAN_PPAP, &kredit.TEMPAT_BEKERJA, &kredit.TGL_AKHIR_BAYAR, &kredit.PIHAK_TERKAIT,
			&kredit.JENIS_JAMINAN, &kredit.NILAI_LEGALITAS, &kredit.RESTRUKTUR_KE, &kredit.TGL_VALID_KOLEK,
			&kredit.TGL_MACET,
		)
		if err != nil {
			return nil, err
		}
		kredits = append(kredits, kredit)
	}

	return kredits, nil
}

// GetByID retrieves a single nominatif_kredit record by ID
func (r *NominatifKreditRepository) GetByID(id int) (*NominatifKredit, error) {
	query := `
		SELECT id, DATADATE, CAB, NOMOR_REKENING, NO_CIF, NAMA_NASABAH, ALAMAT, KODE_KOLEK,
		       JML_HRI_PKK, JML_HARI_BGA, JML_HARI_TUNGGAKAN, KD_PRD, KET_KD_PRD, NOMOR_PERJANJIAN,
		       NO_AKSEP, TGL_PK, TGL_AWAL_FAS, TGL_AKHIR_FAS, TGL_AWAL_AKSEP, TGL_AKH_AKSEP,
		       PLAFOND_AWAL, BAKI_DEBET, LONGGAR_TARIK, BGA, TUNGGAKAN_POKOK, TUNGGAKAN_BUNGA,
		       BGA_JTH_TEMPO, SMP_TGL_CADANG, NILAI_CADANG, ANGSURAN_TOTAL, TGL_PROSES_DENDA,
		       AKUM_DENDA_PKK, AKUM_DENDA_BGA, PRD_AMORT, PRDK_AMORT, FLAG, TGL_AMORT,
		       NILAI_BIAYA_PROVISI, AMORTISASI_PER_PRD, SISA_AMORT_PROV, TAGIH_BIAYA_PROV,
		       NILAI_BIAYA_ADM, AMORT_ADM_PER_PRD, SISA_AMORT_ADM, BYA_ASURANSI, BYA_NOTARIS,
		       PKK_JATEM, BGA_JATEM, REK_BYR_PKK_BGA, SLD_REK_DB, KD_INSTANSI, NM_INSTANSI,
		       REK_BENDAHARA, SFT_KRD, GOL_KRD, JNS_KRD, SKTR_EKNM, ORNTS, NO_HP,
		       POKOK_PINJAMAN, TITIPAN_EFEKTIF, JANGKA_WAKTU, REK_PENCAIRAN, NO_REKENING_LAMA,
		       CIF_LAMA, KODE_GROUP, KET_GROUP, TGL_LAHIR, NIK, NIP, NILAI_BYA_TRANS,
		       AMORT_TRANS_PER_PRD, SISA_AMORT_TRANS, AO, CAB_REK, KELURAHAN, KECAMATAN,
		       CADANGAN_PPAP, TEMPAT_BEKERJA, TGL_AKHIR_BAYAR, PIHAK_TERKAIT, JENIS_JAMINAN,
		       NILAI_LEGALITAS, RESTRUKTUR_KE, TGL_VALID_KOLEK, TGL_MACET
		FROM nominatif_kredit 
		WHERE id = ?`

	var kredit NominatifKredit
	err := r.DB.QueryRow(query, id).Scan(
		&kredit.ID, &kredit.DATADATE, &kredit.CAB, &kredit.NOMOR_REKENING, &kredit.NO_CIF,
		&kredit.NAMA_NASABAH, &kredit.ALAMAT, &kredit.KODE_KOLEK, &kredit.JML_HRI_PKK,
		&kredit.JML_HARI_BGA, &kredit.JML_HARI_TUNGGAKAN, &kredit.KD_PRD, &kredit.KET_KD_PRD,
		&kredit.NOMOR_PERJANJIAN, &kredit.NO_AKSEP, &kredit.TGL_PK, &kredit.TGL_AWAL_FAS,
		&kredit.TGL_AKHIR_FAS, &kredit.TGL_AWAL_AKSEP, &kredit.TGL_AKH_AKSEP, &kredit.PLAFOND_AWAL,
		&kredit.BAKI_DEBET, &kredit.LONGGAR_TARIK, &kredit.BGA, &kredit.TUNGGAKAN_POKOK,
		&kredit.TUNGGAKAN_BUNGA, &kredit.BGA_JTH_TEMPO, &kredit.SMP_TGL_CADANG, &kredit.NILAI_CADANG,
		&kredit.ANGSURAN_TOTAL, &kredit.TGL_PROSES_DENDA, &kredit.AKUM_DENDA_PKK, &kredit.AKUM_DENDA_BGA,
		&kredit.PRD_AMORT, &kredit.PRDK_AMORT, &kredit.FLAG, &kredit.TGL_AMORT, &kredit.NILAI_BIAYA_PROVISI,
		&kredit.AMORTISASI_PER_PRD, &kredit.SISA_AMORT_PROV, &kredit.TAGIH_BIAYA_PROV, &kredit.NILAI_BIAYA_ADM,
		&kredit.AMORT_ADM_PER_PRD, &kredit.SISA_AMORT_ADM, &kredit.BYA_ASURANSI, &kredit.BYA_NOTARIS,
		&kredit.PKK_JATEM, &kredit.BGA_JATEM, &kredit.REK_BYR_PKK_BGA, &kredit.SLD_REK_DB,
		&kredit.KD_INSTANSI, &kredit.NM_INSTANSI, &kredit.REK_BENDAHARA, &kredit.SFT_KRD,
		&kredit.GOL_KRD, &kredit.JNS_KRD, &kredit.SKTR_EKNM, &kredit.ORNTS, &kredit.NO_HP,
		&kredit.POKOK_PINJAMAN, &kredit.TITIPAN_EFEKTIF, &kredit.JANGKA_WAKTU, &kredit.REK_PENCAIRAN,
		&kredit.NO_REKENING_LAMA, &kredit.CIF_LAMA, &kredit.KODE_GROUP, &kredit.KET_GROUP,
		&kredit.TGL_LAHIR, &kredit.NIK, &kredit.NIP, &kredit.NILAI_BYA_TRANS, &kredit.AMORT_TRANS_PER_PRD,
		&kredit.SISA_AMORT_TRANS, &kredit.AO, &kredit.CAB_REK, &kredit.KELURAHAN, &kredit.KECAMATAN,
		&kredit.CADANGAN_PPAP, &kredit.TEMPAT_BEKERJA, &kredit.TGL_AKHIR_BAYAR, &kredit.PIHAK_TERKAIT,
		&kredit.JENIS_JAMINAN, &kredit.NILAI_LEGALITAS, &kredit.RESTRUKTUR_KE, &kredit.TGL_VALID_KOLEK,
		&kredit.TGL_MACET,
	)
	
	if err != nil {
		return nil, err
	}

	return &kredit, nil
}

// GetByNomorRekening retrieves nominatif_kredit records by account number
func (r *NominatifKreditRepository) GetByNomorRekening(nomorRekening string) ([]NominatifKredit, error) {
	query := `
		SELECT id, DATADATE, CAB, NOMOR_REKENING, NO_CIF, NAMA_NASABAH, ALAMAT, KODE_KOLEK,
		       JML_HRI_PKK, JML_HARI_BGA, JML_HARI_TUNGGAKAN, KD_PRD, KET_KD_PRD, NOMOR_PERJANJIAN,
		       NO_AKSEP, TGL_PK, TGL_AWAL_FAS, TGL_AKHIR_FAS, TGL_AWAL_AKSEP, TGL_AKH_AKSEP,
		       PLAFOND_AWAL, BAKI_DEBET, LONGGAR_TARIK, BGA, TUNGGAKAN_POKOK, TUNGGAKAN_BUNGA,
		       BGA_JTH_TEMPO, SMP_TGL_CADANG, NILAI_CADANG, ANGSURAN_TOTAL, TGL_PROSES_DENDA,
		       AKUM_DENDA_PKK, AKUM_DENDA_BGA, PRD_AMORT, PRDK_AMORT, FLAG, TGL_AMORT,
		       NILAI_BIAYA_PROVISI, AMORTISASI_PER_PRD, SISA_AMORT_PROV, TAGIH_BIAYA_PROV,
		       NILAI_BIAYA_ADM, AMORT_ADM_PER_PRD, SISA_AMORT_ADM, BYA_ASURANSI, BYA_NOTARIS,
		       PKK_JATEM, BGA_JATEM, REK_BYR_PKK_BGA, SLD_REK_DB, KD_INSTANSI, NM_INSTANSI,
		       REK_BENDAHARA, SFT_KRD, GOL_KRD, JNS_KRD, SKTR_EKNM, ORNTS, NO_HP,
		       POKOK_PINJAMAN, TITIPAN_EFEKTIF, JANGKA_WAKTU, REK_PENCAIRAN, NO_REKENING_LAMA,
		       CIF_LAMA, KODE_GROUP, KET_GROUP, TGL_LAHIR, NIK, NIP, NILAI_BYA_TRANS,
		       AMORT_TRANS_PER_PRD, SISA_AMORT_TRANS, AO, CAB_REK, KELURAHAN, KECAMATAN,
		       CADANGAN_PPAP, TEMPAT_BEKERJA, TGL_AKHIR_BAYAR, PIHAK_TERKAIT, JENIS_JAMINAN,
		       NILAI_LEGALITAS, RESTRUKTUR_KE, TGL_VALID_KOLEK, TGL_MACET
		FROM nominatif_kredit 
		WHERE NOMOR_REKENING = ?
		ORDER BY id DESC`

	rows, err := r.DB.Query(query, nomorRekening)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kredits []NominatifKredit
	for rows.Next() {
		var kredit NominatifKredit
		err := rows.Scan(
			&kredit.ID, &kredit.DATADATE, &kredit.CAB, &kredit.NOMOR_REKENING, &kredit.NO_CIF,
			&kredit.NAMA_NASABAH, &kredit.ALAMAT, &kredit.KODE_KOLEK, &kredit.JML_HRI_PKK,
			&kredit.JML_HARI_BGA, &kredit.JML_HARI_TUNGGAKAN, &kredit.KD_PRD, &kredit.KET_KD_PRD,
			&kredit.NOMOR_PERJANJIAN, &kredit.NO_AKSEP, &kredit.TGL_PK, &kredit.TGL_AWAL_FAS,
			&kredit.TGL_AKHIR_FAS, &kredit.TGL_AWAL_AKSEP, &kredit.TGL_AKH_AKSEP, &kredit.PLAFOND_AWAL,
			&kredit.BAKI_DEBET, &kredit.LONGGAR_TARIK, &kredit.BGA, &kredit.TUNGGAKAN_POKOK,
			&kredit.TUNGGAKAN_BUNGA, &kredit.BGA_JTH_TEMPO, &kredit.SMP_TGL_CADANG, &kredit.NILAI_CADANG,
			&kredit.ANGSURAN_TOTAL, &kredit.TGL_PROSES_DENDA, &kredit.AKUM_DENDA_PKK, &kredit.AKUM_DENDA_BGA,
			&kredit.PRD_AMORT, &kredit.PRDK_AMORT, &kredit.FLAG, &kredit.TGL_AMORT, &kredit.NILAI_BIAYA_PROVISI,
			&kredit.AMORTISASI_PER_PRD, &kredit.SISA_AMORT_PROV, &kredit.TAGIH_BIAYA_PROV, &kredit.NILAI_BIAYA_ADM,
			&kredit.AMORT_ADM_PER_PRD, &kredit.SISA_AMORT_ADM, &kredit.BYA_ASURANSI, &kredit.BYA_NOTARIS,
			&kredit.PKK_JATEM, &kredit.BGA_JATEM, &kredit.REK_BYR_PKK_BGA, &kredit.SLD_REK_DB,
			&kredit.KD_INSTANSI, &kredit.NM_INSTANSI, &kredit.REK_BENDAHARA, &kredit.SFT_KRD,
			&kredit.GOL_KRD, &kredit.JNS_KRD, &kredit.SKTR_EKNM, &kredit.ORNTS, &kredit.NO_HP,
			&kredit.POKOK_PINJAMAN, &kredit.TITIPAN_EFEKTIF, &kredit.JANGKA_WAKTU, &kredit.REK_PENCAIRAN,
			&kredit.NO_REKENING_LAMA, &kredit.CIF_LAMA, &kredit.KODE_GROUP, &kredit.KET_GROUP,
			&kredit.TGL_LAHIR, &kredit.NIK, &kredit.NIP, &kredit.NILAI_BYA_TRANS, &kredit.AMORT_TRANS_PER_PRD,
			&kredit.SISA_AMORT_TRANS, &kredit.AO, &kredit.CAB_REK, &kredit.KELURAHAN, &kredit.KECAMATAN,
			&kredit.CADANGAN_PPAP, &kredit.TEMPAT_BEKERJA, &kredit.TGL_AKHIR_BAYAR, &kredit.PIHAK_TERKAIT,
			&kredit.JENIS_JAMINAN, &kredit.NILAI_LEGALITAS, &kredit.RESTRUKTUR_KE, &kredit.TGL_VALID_KOLEK,
			&kredit.TGL_MACET,
		)
		if err != nil {
			return nil, err
		}
		kredits = append(kredits, kredit)
	}

	return kredits, nil
}

// Count returns the total number of records
func (r *NominatifKreditRepository) Count() (int, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM nominatif_kredit").Scan(&count)
	return count, err
}

// buildWhereClause builds WHERE clause for filters
func (r *NominatifKreditRepository) buildWhereClause(filters FilterParams) (string, []interface{}) {
	var conditions []string
	var args []interface{}

	if filters.CAB != "" {
		conditions = append(conditions, "CAB = ?")
		args = append(args, filters.CAB)
	}

	if filters.AO != "" {
		conditions = append(conditions, "AO LIKE ?")
		args = append(args, "%"+filters.AO+"%")
	}

	if filters.KET_KD_PRD != "" {
		conditions = append(conditions, "KET_KD_PRD LIKE ?")
		args = append(args, "%"+filters.KET_KD_PRD+"%")
	}

	if filters.TEMPAT_BEKERJA != "" {
		conditions = append(conditions, "TEMPAT_BEKERJA LIKE ?")
		args = append(args, "%"+filters.TEMPAT_BEKERJA+"%")
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	return whereClause, args
}

// GetAllWithFilters retrieves all nominatif_kredit records with filters and pagination
func (r *NominatifKreditRepository) GetAllWithFilters(limit, offset int, filters FilterParams) ([]NominatifKredit, error) {
	whereClause, whereArgs := r.buildWhereClause(filters)
	
	query := `
		SELECT id, DATADATE, CAB, NOMOR_REKENING, NO_CIF, NAMA_NASABAH, ALAMAT, KODE_KOLEK,
		       JML_HRI_PKK, JML_HARI_BGA, JML_HARI_TUNGGAKAN, KD_PRD, KET_KD_PRD, NOMOR_PERJANJIAN,
		       NO_AKSEP, TGL_PK, TGL_AWAL_FAS, TGL_AKHIR_FAS, TGL_AWAL_AKSEP, TGL_AKH_AKSEP,
		       PLAFOND_AWAL, BAKI_DEBET, LONGGAR_TARIK, BGA, TUNGGAKAN_POKOK, TUNGGAKAN_BUNGA,
		       BGA_JTH_TEMPO, SMP_TGL_CADANG, NILAI_CADANG, ANGSURAN_TOTAL, TGL_PROSES_DENDA,
		       AKUM_DENDA_PKK, AKUM_DENDA_BGA, PRD_AMORT, PRDK_AMORT, FLAG, TGL_AMORT,
		       NILAI_BIAYA_PROVISI, AMORTISASI_PER_PRD, SISA_AMORT_PROV, TAGIH_BIAYA_PROV,
		       NILAI_BIAYA_ADM, AMORT_ADM_PER_PRD, SISA_AMORT_ADM, BYA_ASURANSI, BYA_NOTARIS,
		       PKK_JATEM, BGA_JATEM, REK_BYR_PKK_BGA, SLD_REK_DB, KD_INSTANSI, NM_INSTANSI,
		       REK_BENDAHARA, SFT_KRD, GOL_KRD, JNS_KRD, SKTR_EKNM, ORNTS, NO_HP,
		       POKOK_PINJAMAN, TITIPAN_EFEKTIF, JANGKA_WAKTU, REK_PENCAIRAN, NO_REKENING_LAMA,
		       CIF_LAMA, KODE_GROUP, KET_GROUP, TGL_LAHIR, NIK, NIP, NILAI_BYA_TRANS,
		       AMORT_TRANS_PER_PRD, SISA_AMORT_TRANS, AO, CAB_REK, KELURAHAN, KECAMATAN,
		       CADANGAN_PPAP, TEMPAT_BEKERJA, TGL_AKHIR_BAYAR, PIHAK_TERKAIT, JENIS_JAMINAN,
		       NILAI_LEGALITAS, RESTRUKTUR_KE, TGL_VALID_KOLEK, TGL_MACET
		FROM nominatif_kredit ` + whereClause + `
		ORDER BY id DESC 
		LIMIT ? OFFSET ?`

	// Combine filter args with pagination args
	args := append(whereArgs, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kredits []NominatifKredit
	for rows.Next() {
		var kredit NominatifKredit
		err := rows.Scan(
			&kredit.ID, &kredit.DATADATE, &kredit.CAB, &kredit.NOMOR_REKENING, &kredit.NO_CIF,
			&kredit.NAMA_NASABAH, &kredit.ALAMAT, &kredit.KODE_KOLEK, &kredit.JML_HRI_PKK,
			&kredit.JML_HARI_BGA, &kredit.JML_HARI_TUNGGAKAN, &kredit.KD_PRD, &kredit.KET_KD_PRD,
			&kredit.NOMOR_PERJANJIAN, &kredit.NO_AKSEP, &kredit.TGL_PK, &kredit.TGL_AWAL_FAS,
			&kredit.TGL_AKHIR_FAS, &kredit.TGL_AWAL_AKSEP, &kredit.TGL_AKH_AKSEP, &kredit.PLAFOND_AWAL,
			&kredit.BAKI_DEBET, &kredit.LONGGAR_TARIK, &kredit.BGA, &kredit.TUNGGAKAN_POKOK,
			&kredit.TUNGGAKAN_BUNGA, &kredit.BGA_JTH_TEMPO, &kredit.SMP_TGL_CADANG, &kredit.NILAI_CADANG,
			&kredit.ANGSURAN_TOTAL, &kredit.TGL_PROSES_DENDA, &kredit.AKUM_DENDA_PKK, &kredit.AKUM_DENDA_BGA,
			&kredit.PRD_AMORT, &kredit.PRDK_AMORT, &kredit.FLAG, &kredit.TGL_AMORT, &kredit.NILAI_BIAYA_PROVISI,
			&kredit.AMORTISASI_PER_PRD, &kredit.SISA_AMORT_PROV, &kredit.TAGIH_BIAYA_PROV, &kredit.NILAI_BIAYA_ADM,
			&kredit.AMORT_ADM_PER_PRD, &kredit.SISA_AMORT_ADM, &kredit.BYA_ASURANSI, &kredit.BYA_NOTARIS,
			&kredit.PKK_JATEM, &kredit.BGA_JATEM, &kredit.REK_BYR_PKK_BGA, &kredit.SLD_REK_DB,
			&kredit.KD_INSTANSI, &kredit.NM_INSTANSI, &kredit.REK_BENDAHARA, &kredit.SFT_KRD,
			&kredit.GOL_KRD, &kredit.JNS_KRD, &kredit.SKTR_EKNM, &kredit.ORNTS, &kredit.NO_HP,
			&kredit.POKOK_PINJAMAN, &kredit.TITIPAN_EFEKTIF, &kredit.JANGKA_WAKTU, &kredit.REK_PENCAIRAN,
			&kredit.NO_REKENING_LAMA, &kredit.CIF_LAMA, &kredit.KODE_GROUP, &kredit.KET_GROUP,
			&kredit.TGL_LAHIR, &kredit.NIK, &kredit.NIP, &kredit.NILAI_BYA_TRANS, &kredit.AMORT_TRANS_PER_PRD,
			&kredit.SISA_AMORT_TRANS, &kredit.AO, &kredit.CAB_REK, &kredit.KELURAHAN, &kredit.KECAMATAN,
			&kredit.CADANGAN_PPAP, &kredit.TEMPAT_BEKERJA, &kredit.TGL_AKHIR_BAYAR, &kredit.PIHAK_TERKAIT,
			&kredit.JENIS_JAMINAN, &kredit.NILAI_LEGALITAS, &kredit.RESTRUKTUR_KE, &kredit.TGL_VALID_KOLEK,
			&kredit.TGL_MACET,
		)
		if err != nil {
			return nil, err
		}
		kredits = append(kredits, kredit)
	}

	return kredits, nil
}

// CountWithFilters returns the total number of records with filters
func (r *NominatifKreditRepository) CountWithFilters(filters FilterParams) (int, error) {
	whereClause, args := r.buildWhereClause(filters)
	
	query := "SELECT COUNT(*) FROM nominatif_kredit " + whereClause
	
	var count int
	err := r.DB.QueryRow(query, args...).Scan(&count)
	return count, err
}
