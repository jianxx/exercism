class SpaceAge:
    earth_year = 0

    def __init__(self, seconds):
        self.earth_year = float(seconds)/31557600
        pass

    def on_earth(self):
        return round(self.earth_year, 2)
        pass

    def on_mercury(self):
        return round(self.earth_year/0.2408467,2)
        pass

    def on_venus(self):
        return round(self.earth_year/0.61519726,2)
        pass

    def on_mars(self):
        return round(self.earth_year/1.8808158,2)
        pass

    def on_jupiter(self):
        return round(self.earth_year/11.862615,2)
        pass

    def on_saturn(self):
        return round(self.earth_year/29.447498,2)
        pass

    def on_uranus(self):
        return round(self.earth_year/84.016846,2)
        pass

    def on_neptune(self):
        return round(self.earth_year/164.79132,2)
        pass
