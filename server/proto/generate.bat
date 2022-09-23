@echo Off
if %errorlevel% ==0 (
    protoc --go_out . *.proto
    @echo Off
    pause
) else (
  echo "Increasing proto version has error, please to check."
  pause
)