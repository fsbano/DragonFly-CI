package main

import "regexp"
import "os"
import "fmt"
import "bytes"
import "io/ioutil"
import "strings"


func StatementSetVariable(Variable []string, str string) []byte {
  CurrentVariable := strings.Split(Variable[0], "=>")

  VariableName := CurrentVariable[0]
  VariableValue := CurrentVariable[1]

  temp  := strings.Split(str, "\n")
  var buffer bytes.Buffer

  replaceVariable := regexp.MustCompile(fmt.Sprintf("%s.(.*)", strings.TrimSpace(VariableName)))
  for _, element := range temp {
      v := regexp.MustCompile(fmt.Sprintf("^%s", strings.TrimSpace(VariableName)))
      if ( len(v.FindString(element)) > 0 ) {
        r := replaceVariable.ReplaceAllString(element, fmt.Sprintf("%s %s", strings.TrimSpace(VariableName), strings.TrimSpace(VariableValue)))
        buffer.WriteString(r + "\n")
      } else {
        buffer.WriteString(element + "\n")
      }
  }

  return([]byte(strings.TrimSuffix(string(buffer.Bytes()),"\n")))
}

func StatementAbsentVariable(AbsentVariable []string) {
   // fmt.Println(AbsentVariable[0])
}

func StatementCommentVariable(CommentVariable []string) {
   // fmt.Println(CommentVariable[0])
}

func main() {
   dsl, err := os.Open("dsl.cfg")
   if err != nil {
     fmt.Println(err)
   }
   defer dsl.Close()

   cookbook, _ := ioutil.ReadFile("dsl.cfg")
   FindAllStringWith := regexp.MustCompile("With.'.*'").FindAllStringSubmatch(string(cookbook), -1)

   for index, _ := range FindAllStringWith {

     ReFilePath := regexp.MustCompile("'/.*'").FindAllStringSubmatch(strings.Join(FindAllStringWith[index], ""), -1)
     Filename := strings.Replace(strings.Join(ReFilePath[0],""), "'", "", -1 )

     fmt.Println(Filename)

     sshd, _ := ioutil.ReadFile(Filename)
     WithContent := regexp.MustCompile(fmt.Sprintf("%s((\\s)(.+)){1,}",strings.Join(FindAllStringWith[index], "")))
     BlockContent := WithContent.FindString(string(cookbook))
     // ----- Set-Variable

     SetVariable := "Set-Variable.*((\\n)?.\\s+\\w+\\s+=>\\s+\\w+(,)?){1,}"
     FindAllStringSetVariable := regexp.MustCompile(SetVariable)
     FindAllVariable := FindAllStringSetVariable.FindString(BlockContent)
     Variable := regexp.MustCompile("\\w+\\s+=>\\s+\\w+").FindAllStringSubmatch(FindAllVariable, -1)
     for _, element := range Variable {
       sshd = StatementSetVariable(element, string(sshd))
     }

     // ----- Absent-Variable
     // fmt.Println("Absent-Variable:")
     ReAbsentVariable := "Absent-Variable.*((\\n)?.(\\s+)?\\w+(\\s+)?(,)?){1,}\\)"
     FindAllStringAbsent := regexp.MustCompile(ReAbsentVariable)
     FindAllAbsentVariable := FindAllStringAbsent.FindString(BlockContent)
     RemoveAbsentVariableStatement := regexp.MustCompile("Absent-Variable")
     CleanAbsentVariableStatement := RemoveAbsentVariableStatement.ReplaceAllString(FindAllAbsentVariable, "")
     AbsentVariable := regexp.MustCompile("\\w+").FindAllStringSubmatch(CleanAbsentVariableStatement, -1)
     for _, element := range AbsentVariable {
       StatementAbsentVariable(element)
     }

     // ---- Comment-Variable
     // fmt.Println("Comment-Variable:")
     ReCommentVariable := "Comment-Variable.*((\\n)?.(\\s+)?\\w+(\\s+)?(,)?){1,}\\)"
     FindAllStringComment := regexp.MustCompile(ReCommentVariable)
     FindAllCommentVariable := FindAllStringComment.FindString(BlockContent)
     RemoveCommentVariableStatement := regexp.MustCompile("Comment-Variable")
     CleanCommentVariableStatement := RemoveCommentVariableStatement.ReplaceAllString(FindAllCommentVariable, "")
     CommentVariable := regexp.MustCompile("\\w+").FindAllStringSubmatch(CleanCommentVariableStatement, -1)
     for _, element := range CommentVariable {
       StatementCommentVariable(element)
     }

     err = ioutil.WriteFile(Filename, sshd, 0644)
     if err != nil {
       fmt.Println("Failure")
     }
   }

   // fmt.Println(string(sshd))
}
