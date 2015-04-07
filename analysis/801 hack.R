library(plyr)
library(dplyr)
library(lubridate)
library(ggplot2)
path <- "data/vehicle_positions/csv/"
filenames <- list.files(path)

file_size <- function(file, route) {
  dat <- read.csv(paste0(path, file))
  if (!is.null(route)) dat <- subset(dat, route_id == route, select = -route_id)
  object.size(dat)
}

dat_size <- ldply(filenames, file_size, route = 801)
sum(dat_size) / 1024^2

read_dat <- function(file, route = NULL) {
  dat <- read.csv(paste0(path,file), stringsAsFactors=FALSE)
  if (!is.null(route)) dat <- subset(dat, route_id == route, select = -route_id)
  subset(dat, route_id==route, select = -route_id, stringsAsFactors=FALSE)
}

dat <- read_dat(filenames[1], route = 801)
print(object.size(dat), units = "Mb")

raw_dat <- ldply(filenames, read_dat, route = 801)

dat <- within(raw_dat, {
  date <- ymd(substr(timestamp, 1, 10))
  time <- hms(substr(timestamp, 12, 19))
  timestamp <- NULL
  vehicle_id <- factor(vehicle_id)
  trip_id <- factor(trip_id)
})

dim(dat)
print(object.size(dat), units = "Mb")

one_trip <- arrange(subset(dat, trip_id==1400656), 
with(one_trip, table(date))
ggplot(one_trip, aes(x = lat, y = lon, color = factor(date))) + 
  geom_point() + geom_line()
