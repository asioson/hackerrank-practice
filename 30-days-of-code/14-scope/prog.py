class Difference:
    def __init__(self, a):
        self.__elements = a

    def computeDifference(self):
        self.maximumDifference = 0
        if len(self.__elements) > 0:
            maxVal = minVal = self.__elements[0]
            for x in self.__elements:
                if x > maxVal:
                    maxVal = x
                elif x < minVal:
                    minVal = x
            self.maximumDifference = maxVal - minVal

# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
