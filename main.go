package main

import (
	"github.com/lgrees/resy-cli/cmd"
	// "github.com/lgrees/resy-cli/internal/utils/http"
)


func main() {
	cmd.Execute()
	// http.DoBook()
}


// import (
//   "fmt"
//   "strings"
//   "net/http"
//   "io/ioutil"
// )

// func main() {

//   url := "https://api.resy.com/3/book"
//   method := "POST"

//   //under test
//   payload := strings.NewReader("book_token=I2_JNYmRg3Ku9qgmQVJsel16S0XOeCryRPE9y0jRgX75s3kjq5FOTUfOYUqeu_dihXdFZTS2imOX8JrWkc2WRzJUs9MnafDvsYKnYSJufsKeqLenFgJ_LmI_MBcpxPjAXVz5aZXzQb6NHJZ9muEFu2SCTg1yHrYhbl8iLJOiJY5jefKwF5kmNuvsX9sXxESCLU9YPt_FGTq8EZ_uDKt3Gs419FZGhCmJKIsCzQXOBV1VKuSazwZqKYD1bdBPCbJKsSj6E5qxGHfC2A4tcRLPsZvixbLV8aNBGuaqjPCjZDR5cHFlLL7QBfQp8qu4epF0ZivyIKUlk77KhUsY5_m_aQqLuVo1pFrk1W5iqDDWjdiVfCqbmCqtQdfGvmmcu%7CT_10ejny%7CXEqI_A69B_qac%7C1WSO64m%7C13RlElxaWTpWjROATuSmXUZwUBVv53GT8jZbc7ExtyGbGYbxO7%7C5nrTX9_8amIvlASr6xibhdh1KsUZ7ngNFJ1ofZiG_QstJAdyu8RM2hmxEESvkfQmRdr6ZJwSeA_yZA5yGoQhBu50zWNudBxLS5BoLfxL8zQ6IutAHG_TBq2n4nzNJ_VWM%7C%7CycQJbJ_qfsir6KLSpLqFi1bio0a9M7n9Nq3hVNjF1zHMFZuEs4nreMY%7CaZG6HcndrW4Y8XOTJIW%7CXdtI9qPoQK7rPRo8Hqpt4uHQMuWiyC2%7Co-9df4e9afbfcaac7c27c87d528c9ab4509064253fdb28d629eb9c1ed2&source_id=resy.com-venue-details&struct_payment_method=%7B%22id%22%3A23574202%7D")

//   //verified official
// //   payload := strings.NewReader("book_token=RBsZFQd9BqAzYtAUTmeusCBAhomZpWwqz1iBUwb46ujh19QB4NrJXgebCy%7CwoYrtCDMMnoGjyYLw4D5wtOorm4sI_p8lgsQANNEK9%7C2bdI7NYq2cx5_jB25qYPX0Q7AHfAGHmn1GHFJnJt2ANxLC3tjBnwlzWcArGsnuSI4MZnVhcF4bq4WhgRoySxYM_lJwPn8N90sIyp4jf0oTs12d_uzRNm7a%7CYhZlhP2E1CGc%7CfNSzzLVMB%7CwEbTlk7QkVvjKt1BaVRdhkmQuag3NAXqqY_PUg6W9UiCukSW%7CFJpk0TlQBmad3rnV5bMcXuv1FTQK0PVLgE_T9MdVuQDsS6suyOH3QC_pFUa6KHPVdu0tUWSJTNti0mN7GBNpFaBI%7CetAJRQRPpF_823xwRKsOYAQWhqAZJ%7CUCWaPiIXgcNsRlle_wtszSZoUdKRk69fThi_6ximnnJ%7CdhmwIZahJ12tBlkO8M4kRhEoe9SjwmvqG9KsVl4dC0tL%7C_6k53DTc8AxKgGYCYGJkNl1FXk60Jb9xx%7CrKwE0NiVrLpLyTitso%7CrCSw%7C9Ohow2TwB8GCDaN40hnWK5SayMy1N0SMIk30IVhyEU1hh%7CXFp7mdd%7C3rKlyZx03dnB_Kp6zZk0p_0IjX2rYGK2MrQWeIPndrWZnIXOe7vzHOh%7CIIlV1y2QkRImaZbbAdbp_1yoi8kr5RmYFro-7d0b0259358e4eb9c7199edbc4eb305a0b83d483d9300633d915ff8f&struct_payment_method=%7B%22id%22%3A23574202%7D&source_id=resy.com-venue-details")

//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, payload)

//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   req.Header.Add("accept", "application/json, text/plain, */*")
//   req.Header.Add("authorization", "ResyAPI api_key=\"VbWk7s3L4KiK5fzlO7JD3Q5EYolJI7n5\"")
//   req.Header.Add("cache-control", "no-cache")
//   req.Header.Add("content-type", "application/x-www-form-urlencoded")
//   req.Header.Add("origin", "https://widgets.resy.com")
//   req.Header.Add("referer", "https://widgets.resy.com/")
//   req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36")
//   req.Header.Add("x-origin", "https://widgets.resy.com")
//   req.Header.Add("x-resy-auth-token", "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE3MTY4MTQ4NTgsInVpZCI6NDUzNjQzLCJndCI6ImNvbnN1bWVyIiwiZ3MiOltdLCJsYW5nIjoiZW4tdXMiLCJleHRyYSI6eyJndWVzdF9pZCI6MTAzMjY5NX19.AaffzkFn88vD6aWrEws27NtR4PckIGSYb9oJwN3SkeSzcJhBV6YRuwl-CskZEyJMzzc6RarFPbUlXnzkp7vg7qXEAUnYNwnQ9-0PKsQNhNfNwTQzbG9jKmKIczmFPGyILw7eqi0JL77Y_FVMilmRtSRC5vX8h4x74AvTCsiZxdssTiTU")
//   req.Header.Add("x-resy-universal-auth", "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NiJ9.eyJleHAiOjE3MTY4MTQ4NTgsInVpZCI6NDUzNjQzLCJndCI6ImNvbnN1bWVyIiwiZ3MiOltdLCJsYW5nIjoiZW4tdXMiLCJleHRyYSI6eyJndWVzdF9pZCI6MTAzMjY5NX19.AaffzkFn88vD6aWrEws27NtR4PckIGSYb9oJwN3SkeSzcJhBV6YRuwl-CskZEyJMzzc6RarFPbUlXnzkp7vg7qXEAUnYNwnQ9-0PKsQNhNfNwTQzbG9jKmKIczmFPGyILw7eqi0JL77Y_FVMilmRtSRC5vX8h4x74AvTCsiZxdssTiTU")

//   res, err := client.Do(req)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer res.Body.Close()

//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   fmt.Println(string(body))
// }