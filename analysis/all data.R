library(plyr)
library(dplyr)
library(lubridate)
library(ggplot2)
path <- "data/vehicle_positions/csv/"
filenames <- list.files(path)

file_size <- function(file, route) {
  dat <- read.csv(paste0(path, file), stringsAsFactors=FALSE)
  if (!missing(route)) dat <- subset(dat, route_id == route, select = -route_id)
  object.size(dat)
}

dat_size <- ldply(filenames, file_size)
sum(dat_size) / 1024^2

read_dat <- function(file, route = NULL) {
  dat <- read.csv(paste0(path,file), stringsAsFactors=FALSE)
  if (!is.null(route)) dat <- subset(dat, route_id == route, select = -route_id)
  dat
}

raw_dat <- ldply(filenames, read_dat)
summary(raw_dat)
table(raw_dat$route_id)

ggplot(subset(raw_dat, lat!=0 & lon != 0), 
       aes(x = lat, y = lon, color = factor(route_id))) + 
  geom_point() + 
  theme_minimal()