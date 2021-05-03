
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
 bolsterNumber     text NOT NULL,
 firstExtReqWeight text NOT NULL,
 solidLeadPI       boolean NOT NULL,
 solidDiePI        boolean NOT NULL,
 solidBacker       boolean NOT NULL,
 portholeDie       boolean NOT NULL,
 portholeMandrel   boolean NOT NULL,
 portholeBacker    boolean NOT NULL,
 description       text NOT NULL,
 size              text NOT NULL,
 kgs               text NOT NULL,
 sup               text NOT NULL,
 price             text NOT NULL,
 remarks           text NULL,
 dieNumber         text NOT NULL,
 cavNumber         numeric NOT NULL,
 companyName       text NOT NULL,
 email             text NOT NULL,
 address           text NOT NULL,
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



