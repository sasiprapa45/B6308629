import { DoctorInterface } from "./IDoctor";
import { PatientTypeInterface } from "./IPatientType";
import { SymptomsInterface } from "./ISymptoms";

export interface PatientsInterface {
  ID: number;

  PatientsName: string;

  Age: number;

  DateAdmit: Date ;

  DoctorID: number;
  Doctor: DoctorInterface;

  PatientTypeID: number;
  PatientType: PatientTypeInterface;
  
  SymptomsID: number;
  Symptoms: SymptomsInterface; 
}