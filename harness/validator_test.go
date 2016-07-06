package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	Nias2 "github.com/nsip/nias2/lib"
	"menteslibres.net/gosexy/rest"
)

var customClient *rest.Client

/*
func TestSexMissingMandatory(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsMissingMandatorySex.csv", "Sex", "Sex is required")
}

func TestSexInvalid(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidSex.csv", "Sex", "Sex must be one of the following")
}

func TestYearLevelPrep(t *testing.T) {
	test_harness(t, "../unit_test_files/1students1YearLevelPrep.csv", "BirthDate/TestLevel", "Year level supplied is P, does not match expected test level")
}

func TestYearLevelF(t *testing.T) {
	test_harness(t, "../unit_test_files/1students2YearLevelF.csv", "BirthDate/YearLevel", "Student Year Level (yr F) does not match year level derived from BirthDate")
}

func TestFutureBirthdate(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsFutureBirthDates.csv", "BirthDate/YearLevel", "Year Level calculated from BirthDate does not fall within expected NAPLAN year level ranges")
}

func TestMissingParent2LOTE(t *testing.T) {
	test_harness(t, "../unit_test_files/1students2MissingParent2LOTE.csv", "Parent2LOTE", "Parent2LOTE is required")
}

func TestACARAIDandStateBlank(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsACARAIDandStateBlank.csv", "ASLSchoolId", "ASLSchoolId is required")
}

func TestBirthdateYearLevel(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsBirthdateYearLevel.csv", "BirthDate/YearLevel/TestLevel", "does not match year level derived from BirthDate")
}

func TestACARAIDandStateMismatch(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsACARAIDandStateMismatch.csv", "ASLSchoolID", "is a valid ID, but not for")
}

func TestMissingSurname(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsMissingSurname.csv", "FamilyName", "FamilyName is required")
}

func TestEmptySurname(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsEmptySurname.csv", "FamilyName", "FamilyName is required")
}

func TestInvalidVisaClass(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidVisaSubClass.csv", "VisaCode", "is not one of known values from")
}

func TestMalformedPSI(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentMalformedPlatformStudentID.csv", "PlatformId", "PlatformId is not in correct format")
}

func TestCommaAddressField(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsCommaAddressField.csv", "", "")
}

func TestDupGivenLastNameDOBDiffACARAId(t *testing.T) {
	test_harness(t, "../unit_test_files/2studentsDupGivenLastNameDOBDiffACARAId.csv", "", "")
}

func TestDupGivenLastNameDOBCARAId(t *testing.T) {
	test_harness(t, "../unit_test_files/2studentsDupGivenLastNameDOBSchool.csv", "Multiple (see description)", "otential duplicate of record")
}

func TestDuplicateStudentOneSchool(t *testing.T) {
	test_harness(t, "../unit_test_files/2studentsDuplicateStudentOneSchool.csv", "LocalID/ASL ID", "LocalID (Student) and ASL ID (School) are potential duplicate of record")
}

func TestExceedCharLengthsSurname(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsExceedCharLengthsSurname.csv", "FamilyName", "String length must be less than or equal to 40")
}

func TestExceedCharLengthsAddress(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsExceedCharLengthsAddress.csv", "AddressLine1", "String length must be less than or equal to 40")
}

func TestExceedCharLengthsGivenName(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsExceedCharLengthsGivenName.csv", "GivenName", "String length must be less than or equal to 40")
}

func TestExceedLengthHomeGrp(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsExceedLengthHomeGrp.csv", "Homegroup", "String length must be less than or equal to 10")
}

func TestInvalidAcaraId(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidAcaraId.csv", "ASLSchoolID", "not found in ASL list of valid IDs")
}

func TestInvalidCountryCodes(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidCountryCodes.csv", "CountryOfBirth", "Country Code is not one of SACC 1269.0 codeset")
}

func TestInvalidDateFormat(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidDateFormat.csv", "BirthDate", "Date provided does not parse correctly for yyyy-mm-dd")
}

func TestInvalidLanguageCodes(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidLanguageCodes.csv", "StudentLOTE", "Language Code is not one of ASCL 1267.0 codeset")
}


func TestInvalidValuesLBOTE(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesLBOTE.csv", "LBOTE", "LBOTE must be one of the following")
}

func TestInvalidValuesOfflineDelivery(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesOfflineDelivery.csv", "OfflineDelivery", "OfflineDelivery must be one of the following")
}

func TestInvalidValuesMainSchoolFlag(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesMainSchoolFlag.csv", "MainSchoolFlag", "MainSchoolFlag must be one of the following")
}

func TestInvalidValuesParent1LOTE(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidParent1LOTE.csv", "Parent1LOTE", "Language Code is not one of ASCL 1267.0 codeset")
}

func TestInvalidInvalidValuesFFPOS(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesFFPOS.csv", "FFPOS", "FFPOS must be one of the following")
}

func TestInvalidValuesParent2Occupation(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesParent2Occupation.csv", "Parent2Occupation", "Parent2Occupation must be one of the following")
}

func TestInvalidValuesParent2NonSchoolEducation(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesParent2NonSchoolEducation.csv", "Parent2NonSchoolEducation", "Parent2NonSchoolEducation must be one of the following")
}

func TestInvalidValuesParent2SchoolEducation(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesParent2SchoolEducation.csv", "Parent2SchoolEducation", "Parent2SchoolEducation must be one of the following")
}

func TestInvalidValuesHomeSchooledStudent(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesHomeSchooledStudent.csv", "HomeSchooledStudent", "HomeSchooledStudent must be one of the following")
}

func TestInvalidValuesYearLevel(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidValuesYearLevel.csv", "YearLevel", "YearLevel must be one of the following")
}

*/

func TestInvalidVisaClass(t *testing.T) {
	test_harness(t, "../unit_test_files/1studentsInvalidVisaSubClass.csv", "VisaCode", "is not one of known values from")
}

/* if errfield is nil, we expect test to pass */
func test_harness(t *testing.T, filename string, errfield string, errdescription string) {
	var f *os.File
	var err error
	bytebuf := []byte{}
	dat := []map[string]string{}

	if f, err = os.Open(filename); err != nil {
		t.Fatalf("Error %s", err)
	}
	defer f.Close()
	files := rest.FileMap{
		"validationFile": []rest.File{{
			Name:   path.Base(f.Name()),
			Reader: f},
		},
	}
	requestVariables := url.Values{"name": {path.Base(f.Name())}}
	msg, err := rest.NewMultipartMessage(requestVariables, files)
	if err != nil {
		t.Fatalf("Error %s", err)
	}
	dst := map[string]interface{}{}
	if err = customClient.PostMultipart(&dst, "/naplan/reg/validate", msg); err != nil {
		t.Fatalf("Error %s", err)
	}
	txid := dst["TxID"].(string)
	time.Sleep(1 * time.Second)
	if err = customClient.Get(&bytebuf, "/naplan/reg/results/"+txid, nil); err != nil {
		t.Fatalf("Error %s", err)
	}
	// we are getting back a JSON array
	if err = json.Unmarshal(bytebuf, &dat); err != nil {
		t.Fatalf("Error %s", err)
	}
	log.Println(dat)
	if errfield == "" {
		if len(dat) > 0 {
			t.Fatalf("Expected no error, got error in %s: %s", dat[0]["errfield"], dat[0]["description"])
		}
	} else {
		if len(dat) < 1 {
			t.Fatalf("Expected error field %s, got no error", errfield)
		} else {
			if dat[0]["errField"] != errfield {
				t.Fatalf("Expected error field %s, got field %s", errfield, dat[0]["errField"])
			}
			if !strings.Contains(dat[0]["description"], errdescription) {
				t.Fatalf("Expected error description %s, got description %s", errdescription, dat[0]["description"])
			}
		}
	}
}

func TestMain(m *testing.M) {
	customClient, _ = rest.New("http://localhost:" + Nias2.NiasConfig.WebServerPort + "/")
	os.Exit(m.Run())
}
