resource "alicloud_slb_listener" "tcp" {
  load_balancer_id          = "lb-bp1ahgwxc931hdp1os7ce"
  backend_port              = 80
  frontend_port             = 80
  protocol                  = "tcp"
  bandwidth                 = 10
  sticky_session            = "on"
  sticky_session_type       = "insert"
  cookie_timeout            = 86400
  cookie                    = "testslblistenercookie"
  health_check              = "on"
  health_check_connect_port = 80
  health_check_type = "tcp"
}