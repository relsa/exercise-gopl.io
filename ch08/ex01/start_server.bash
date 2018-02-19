go build ./clock.go
TZ=US/Eastern    ./clock -port 8010 &
TZ=Asia/Tokyo    ./clock -port 8020 &
TZ=Europe/London ./clock -port 8030 &
