program project1;

{$mode objfpc}{$H+}

uses
  {$IFDEF UNIX}{$IFDEF UseCThreads}
  cthreads,
  {$ENDIF}{$ENDIF}
  Interfaces, // this includes the LCL widgetset
  Forms, unit2, runtimetypeinfocontrols
  { you can add units after this };

{$R *.res}

begin
  RequireDerivedFormResource:=True;
  Application.Title:='PycByte';
  Application.Scaled:=True;
  Application.Initialize;
  Application.CreateForm(TMain, Main);
  Application.Run;
end.

