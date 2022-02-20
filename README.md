### 減量をする際の最適なPFCバランスを計算するCLIを作成
height：身長
weight：体重
age：年齢
  active_level：1日の活動量
* ほぼ運動しない (基礎代謝 × 1.2)
* 軽い運動 (基礎代謝 × 1.3) 
* 中程度の運動 (基礎代謝 × 1.5)
* 激しい運動 (基礎代謝 × 1.7)
* 非常に激しい (基礎代謝 × 1.9)

`go run main.go --height 187 --weight 82 --age 28 --active_level 1.2`  
上記コマンドを実行すると下記情報が出力される。  
【あなたの1日基礎代謝量】  
 1925  
【減量時の1日最適摂取カロリー】  
 2310  
【1日のPFCバランス】  
 タンパク質: 577  
 脂質: 346  
 炭水化物: 1386  
