package handler

import "github.com/gin-gonic/gin"

func (h *Handler) GetBusSchedule(g *gin.Context) {
	// // get bus schedule
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*transportClient).GetBusSchedule(ctx, &pbTransport.RequestBusSchedule{BusNumber: 5})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Schedules)
}

func (h *Handler) TrackBusLocation(g *gin.Context) {
	// // get current location of bus
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*transportClient).TrackBusLocation(ctx, &pbTransport.RequestBusLocation{BusNumber: 5})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.StationName)
}

func (h *Handler) ReportTrafficJam(g *gin.Context) {
	// // give report for traffic Jam
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// resp, err := (*transportClient).ReportTrafficJam(ctx, &pbTransport.RequestTrafficJam{Report: "chiiiiiiiii"})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Success)
}
