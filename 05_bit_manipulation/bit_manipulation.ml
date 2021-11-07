open Stdint.Int32

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
10000000000 10011

11110000000 = -1 << (6 + 1)
00000000011 = (1 << 2) - 1

(10000000000 & (11110000000 | 00000000011)) | 1001100
(10000000000 & 11110000011) | 1001100
10001001100
 *)
let insertion n m i j =
  let lmask = shift_left (-1l) (Int.add j 1) in (* 11110000000 *)
  let rmask = shift_left 1l i - 1l in (* 00000000011 *)
  let mask = logor lmask rmask in (* 11110000011 *)
  let masked_n = logand n mask in
  let shifted_m = shift_left m i in
  logor masked_n shifted_m

let () =
  assert (Int32.equal (insertion 0b10000000000l 0b10011l 2 6) 0b10001001100l);
  assert (Int32.equal (insertion 0b11111111111l 0b00110l 3 7) 0b11100110111l);
