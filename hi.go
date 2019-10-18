package hi

import (
   "errors"
   "fmt"
)

func Hi(name, lang string) (string, error) {
   switch lang {
   case "en":
      return fmt.Sprintf("Welcome, %s!", name), nil
   case "tw":
      return fmt.Sprintf("嗨, %s!", name), nil
   case "fr":
      return fmt.Sprintf("Bonjour, %s!", name), nil
   default:
      return "", errors.New("unknown language")
   }
}
