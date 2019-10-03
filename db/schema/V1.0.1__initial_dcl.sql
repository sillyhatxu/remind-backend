CREATE USER sillyhat_remind identified BY 'sillyhat_remind';
CREATE SCHEMA `sillyhat_remind` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin ;
GRANT ALL ON sillyhat_remind.* TO sillyhat_remind;
GRANT ALL ON sillyhat_remind.* TO sillyhat;