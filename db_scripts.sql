
-- ************************************** dieMasterRegistry
CREATE TABLE dieMasterRegistry
(
 lotNumber  numeric NOT NULL,
 dieNumber  text NOT NULL,
 nosCavity  numeric NOT NULL,
 type       text NOT NULL,
 bolsterNo  text NOT NULL,
 supplier   text NOT NULL,
 dieSize    text NOT NULL,
 orderDate  date NOT NULL,
 landedDate date NOT NULL,
 price      numeric NOT NULL,
 CONSTRAINT PK_diemasterregistry PRIMARY KEY ( lotNumber, dieNumber )
);


-- ************************************** bolsterMasterRegistry

CREATE TABLE bolsterMasterRegistry
(
 lotNumber     numeric NOT NULL,
 bolsterNumber text NOT NULL,
 nosCavity     numeric NOT NULL,
 type          text NOT NULL,
 supplier      text NOT NULL,
 dieSize       text NOT NULL,
 orderDate     date NOT NULL,
 landedDate    date NOT NULL,
 price         numeric NOT NULL,
 bolsterDesign text NOT NULL,
 CONSTRAINT PK_diemasterregistry_clone PRIMARY KEY ( lotNumber, bolsterNumber )
);


-- ************************************** users

CREATE TABLE users
(
 username  text NOT NULL,
 password text NOT NULL,
 role     text NOT NULL,
 CONSTRAINT PK_users PRIMARY KEY ( username )
);


-- ************************************** dieOrder

CREATE TABLE dieOrder
(
 lotNumber         numeric NOT NULL,
 sl                numeric NOT NULL,
 bolsterNumber     text ,
 firstExtReqWeight text ,
 solidLeadPI       boolean ,
 solidDiePI        boolean ,
 solidBacker       boolean ,
 portholeDie       boolean ,
 portholeMandrel   boolean ,
 portholeBacker    boolean ,
 description       text ,
 size              text ,
 kgs               text ,
 sup               text ,
 price             text ,
 remarks           text ,
 dieNumber         text ,
 cavNumber         numeric ,
 companyName       text ,
 email             text ,
 address           text ,
 CONSTRAINT PK_dietable PRIMARY KEY ( lotNumber, sl )
);

-- ************************************** dieInspectionReport

CREATE TABLE dieInspectionReport
(
 lotNumber        numeric NOT NULL,
 slNumber         numeric NOT NULL,
 receivedDate     date NOT NULL,
 dieNumber        text NOT NULL,
 cav              numeric NOT NULL,
 dieDetailsMF     text NOT NULL,
 dieDetailsDie    text NOT NULL,
 dieDetailsBacker text NOT NULL,
 price            text NOT NULL,
 hardnessMF       text NOT NULL,
 hardnessDie      text NOT NULL,
 hardnessBacker   text NOT NULL,
 CONSTRAINT PK_dieinspectionreport PRIMARY KEY ( lotNumber, slNumber )
);



