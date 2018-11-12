/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `scores` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ResultAlias` varchar(40) NOT NULL,
  `NBodyTime` float NOT NULL,
  `NBodyScore` float NOT NULL,
  `PiDigitsTime` float NOT NULL,
  `PiDigitsScore` float NOT NULL,
  `MandelbrotTime` float NOT NULL,
  `MandelbrotScore` float NOT NULL,
  `SpectralNormTime` float NOT NULL,
  `SpectralNormScore` float NOT NULL,
  `BinaryTreesTime` float NOT NULL,
  `BinaryTreesScore` float NOT NULL,
  `TotalTime` float NOT NULL,
  `TotalScore` float NOT NULL,
  `Overclocked` int(1) NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;
