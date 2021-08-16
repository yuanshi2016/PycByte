unit Unit2;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, MaskEdit,
  Buttons, RTTICtrls, SynEdit;

type

  { TMain }

  TMain = class(TForm)
    PycSelectButton: TButton;
    SaveSelectButton: TButton;
    RunButton: TButton;
    LogEdit: TEdit;
    PycPathEdit: TMaskEdit;
    SavePathEdit: TMaskEdit;
    SavePathLabel: TLabel;
    PycPathLabel: TLabel;
    procedure FormCreate(Sender: TObject);
    procedure PycSelectButtonClick(Sender: TObject);
    procedure RunButtonClick(Sender: TObject);
    procedure SaveSelectButtonClick(Sender: TObject);
  private

  public

  end;

var
  Main: TMain;

implementation

{$R *.lfm}

{ TMain }

procedure TMain.FormCreate(Sender: TObject);
begin

end;

procedure TMain.PycSelectButtonClick(Sender: TObject);
begin

end;

procedure TMain.RunButtonClick(Sender: TObject);
begin

end;

procedure TMain.SaveSelectButtonClick(Sender: TObject);
begin

end;

end.

