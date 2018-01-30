defmodule BoletoBarCode do
  defstruct value: "00000000000000000000000000000000000000000000"
end

defmodule BoletoWrittenLineCode do
  defstruct value: "000000000000000000000000000000000000000000000000"
end

defmodule CampoLivreBoleto do
  defstruct bank_branch_code: "0000",
    wallet_code: "00",
    nosso_numero: "00000000000",
    assignor_account: "0000000"
end

defmodule BoletoInfo do
  defstruct bank_code: "000", 
    currency_code: "0", 
    expiration_factor: "0000", 
    nominal_value: "0000000000",
    campo_livre: %CampoLivreBoleto{}
end

defmodule GDACode do
  defstruct value: "00000000000000000000000000000000000000000000"
end

defmodule CampoLivreGDA do
  defstruct expiration_date: "AAAAMMDD",
    other_data: "0000000000000"
end

defmodule GDAInfo do
  defstruct product_id: "0",
    segment_id: "0",
    real_value_id: "0",
    gda_value: "00000000000",
    cnpj_mf: "00000000",
    campo_livre: %CampoLivreGDA{}
end

defmodule GDAPayment do
  def from_bar_code(value \\ %GDACode{}.value)
  def to_bar_code(information \\ %GDAInfo{})

  @spec from_bar_code(String.t) :: {:ok, GDAInfo.t} | :error
  def from_bar_code(value) do
    
  end
  @spec to_bar_code(GDAInfo.t) :: {:ok, String.t} | :error # Será?
  def to_bar_code(information) do
    
  end
end

defmodule BoletoPayment do
  @spec from_written_line(String.t) :: {:ok, BoletoInfo.t} | :error
  @spec from_bar_code(String.t) :: {:ok, BoletoInfo.t} | :error
  @spec to_written_line(BoletoInfo.t) :: {:ok, String.t} | :error
  @spec to_bar_code(BoletoInfo.t) :: {:ok, String.t} | :error
end