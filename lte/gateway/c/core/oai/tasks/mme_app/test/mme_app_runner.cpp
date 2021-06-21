

int main(int argc, char* argv[]) {
  srand(time(NULL));

  CHECK_INIT_RETURN(OAILOG_INIT(
      MME_CONFIG_STRING_MME_CONFIG, OAILOG_LEVEL_DEBUG, MAX_LOG_PROTOS));
  CHECK_INIT_RETURN(shared_log_init(MAX_LOG_PROTOS));
  CHECK_INIT_RETURN(itti_init(
      TASK_MAX, THREAD_MAX, MESSAGES_ID_MAX, tasks_info, messages_info, NULL,
      NULL));

  CHECK_INIT_RETURN(timer_init());
  // Could not be launched before ITTI initialization
  shared_log_itti_connect();
  OAILOG_ITTI_CONNECT();
  mme_app_init(mme_config_);
}
