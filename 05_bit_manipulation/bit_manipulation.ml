(* 
bit manipulation

examples
- 0110 + 0010 = 1000
- 0011 * 0101 = 1111
- 0110 + 0110 = 1100 // 0110 * 2 = 0110 << 1
- 0011 + 0010 = 0101
- 0011 * 0011 = 1001
- 0100 * 0011 = 1100 // 0011 << 2
- 0110 - 0011 = 0011
- 1101 >> 2   = 0011
- 1101 ^ (~1101) = 1111 // a^(~a) = 1111
- 1000 - 0110 = 0010
- 1101 ^ 0101 = 1000
- 1011 & (~0 << 2) = 1000
 *)


(* 
5.2 Insertion
10000000000 10011

11110000000 = -1 << (6 + 1)
00000000011 = (1 << 2) - 1

(10000000000 & (11110000000 | 00000000011)) | 1001100
(10000000000 & 11110000011) | 1001100
10001001100
 *)
let insertion n m i j =
  let mask = (-1 lsl (j + 1)) lor ((1 lsl i) - 1) in (* 11110000011 *)
  (n land mask) lor (m lsl i)

let () =
  assert ((insertion 0b10000000000 0b10011 2 6) = 0b10001001100);
  assert ((insertion 0b11111111111 0b00110 3 7) = 0b11100110111);
;;

(* 
5.3 Binary to String
 *)  
let to_string_bin d = 
  if d > 1. || d < 0. then "ERROR" else
  let rec f b i d = 
    if i = 32 then if d = 0. then b else "ERROR" else
    let threshold = Float.(pow 2. (of_int (-i))) in
    if d >= threshold then f (b ^ "1") (i + 1) (d -. threshold) else f (b ^ "0") (i + 1) d
  in
  f "" 1 d

let () =
  let test d =
    let got = to_string_bin d in
    print_endline got;
    let gotf = Int32.(to_float (of_string ("0b" ^ got))) in
    let wantf = d *. Float.pow 2. 31. in
    Printf.printf "got: %f, want: %f\n" gotf wantf;
    assert(gotf = wantf)
  in
  test 0.5;
  test 0.25;
;;



