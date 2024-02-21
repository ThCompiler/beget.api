## 0.0.1

## Added

### Library

#### Core

* Added error types, formats and statuses (core/info section).
* The structure of the response description from the api service has been added.
* Added generalized execution of api requests and getting results.
* Added the `ApiMethod` interface describing the requirements for the api method.

#### Api

* Added a basic method describing the api request.
* Added methods for interacting with DNS: *getData*, *changeRecords*.
* Implemented a record creator for requesting dns record changes via the *changeRecords* method.

### Repository

* Added a description of the README in Russian and English.
* CI operation is configured.
* Added an example of using the library in another repository.

## 0.0.2

## Fixed

### Library

#### Core

* Fixed return ***nil*** error from method *Get* for *BegetResponse* and *Answer*.

### Repository

* Fixed the path to documentation in README.
* Fixed the path to get a package via the `go get` command in the README.

## Changed

### Library

#### Core

* Renamed methods *GetResult* and *MustGetResult* of *BegetResponse* to *GetAnswer* and *MustGetAnswer*.
* The package structure has been redesigned.

#### Api

* Made exportable *settableRecords* interface in the **dns** package.
* Changed type of **records** field in *CallChangeRecords* method to *SettableRecords* in the **dns** package.
* Deleted unnecessary *SetRecords*, *SetBasicRecords*, *SetNsRecords*, *SetCNameRecords* functions in the **dns** package.
* Deleted unnecessary *SettingRecords* type in the **dns** package.

## Added

### Library

#### Core

* Added golang documentation.

#### Api

* Added golang documentation.