terraform {
  source = "../..//apartments"
}

dependencies {
  paths = ["../first-phase", "../second-phase"]
}

inputs = {
  name                  = "block a"
  no_of_1bhk_apartments = 0
  no_of_2bhk_apartments = 0
  no_of_3bhk_apartments = 1
}