library(dplyr)
library(ggplot2)

stops <- read.csv("GTFS/stops.txt")
stop_times <- read.csv("GTFS/stop_times.txt")
trips <- read.csv("GTFS/trips.txt")
routes <- read.csv("GTFS/routes.txt")
shapes <- read.csv("GTFS/shapes.txt")

trips_801 <- subset(trips, route_id==801)
subset(stop)

table(shapes$shape_id)
subset(stop_times, trip_id==1370220)

a_trip <- merge(subset(stop_times, trip_id==1370220), stops, by = "stop_id", all.x = TRUE)
head(a_trip)

qplot(stop_lat, stop_lon, data = a_trip, geom = "line")


ggplot(shapes, aes(shape_pt_lon, shape_pt_lat, col = factor(shape_id))) + 
  geom_line() + 
  theme_minimal()
 