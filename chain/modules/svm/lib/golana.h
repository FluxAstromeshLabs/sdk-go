/*
 * golana is a C interface for Solana SeaLevel VM.
*/

#pragma once

/* Generated with cbindgen:0.26.0 */

/* Warning, this file is autogenerated by cbindgen. Don't modify this manually. */

#include <stdarg.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct c_compute_budget c_compute_budget;

typedef struct c_instruction_account c_instruction_account;

typedef struct c_ix_info c_ix_info;

typedef struct c_pubkeys c_pubkeys;

typedef struct c_result c_result;

typedef struct c_transaction_account c_transaction_account;

typedef struct c_tx_callback c_tx_callback;

typedef struct {
  size_t len;
  const uint8_t *data;
} bytes;

typedef c_compute_budget *(*getComputeBudgetFn)(void *caller, uint64_t tx_id);

typedef c_pubkeys *(*getPubkeysFn)(void *caller, uint64_t tx_id);

typedef uint64_t (*getIxLenFn)(void *caller, uint64_t tx_id);

typedef c_ix_info *(*getIxInfoFn)(void *caller, uint64_t tx_id, uint64_t ix_id);

typedef c_transaction_account *(*getAccountSharedDataFn)(void *caller, uint8_t *pubkey);

typedef bool (*setAccountSharedDataFn)(void *caller, c_transaction_account *account);

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

c_result *execute(c_tx_callback *cb, uint64_t tx_id, uint64_t *total_unit_consumed);

bytes get_builtins_program_keys(void);

c_tx_callback *tx_callback_create(getComputeBudgetFn get_compute_budget_fn,
                                  getPubkeysFn get_pubkeys_fn,
                                  getIxLenFn get_ix_len_fn,
                                  getIxInfoFn get_ix_info_fn,
                                  getAccountSharedDataFn get_account_shared_data_fn,
                                  setAccountSharedDataFn set_account_shared_data_fn);

void tx_callback_free(c_tx_callback *c_tx_callback);

void bytes_free(bytes b);

c_transaction_account *transaction_account_create(const uint8_t *pubkey_ptr,
                                                  const uint8_t *owner_ptr,
                                                  uint64_t lamports,
                                                  const uint8_t *data_ptr,
                                                  size_t data_len,
                                                  bool executable,
                                                  uint64_t rent_epoch);

void transaction_account_debug(const c_transaction_account *c);

const uint8_t *transaction_account_get_pubkey(c_transaction_account *a);

uint64_t transaction_account_get_lamports(c_transaction_account *a);

bytes transaction_account_get_data(c_transaction_account *a);

const uint8_t *transaction_account_get_owner(c_transaction_account *a);

bool transaction_account_get_executable(c_transaction_account *a);

uint64_t transaction_account_rent_epoch(c_transaction_account *a);

c_compute_budget *compute_budget_create(uint64_t compute_unit_limit,
                                        size_t max_instruction_trace_length,
                                        size_t max_invoke_stack_height);

c_pubkeys *pubkeys_create(const uint8_t *const *ptr, size_t len);

c_instruction_account *instruction_account_create(uint16_t index_in_transaction,
                                                  uint16_t index_in_caller,
                                                  uint16_t index_in_callee,
                                                  bool is_signer,
                                                  bool is_writable);

c_ix_info *ix_info_create(c_instruction_account **accounts_ptr,
                          size_t accounts_len,
                          uint16_t *program_accounts_ptr,
                          size_t program_accounts_len,
                          uint8_t *data_ptr,
                          size_t data_len);

void c_result_free(c_result *wrapper);

const char *c_result_error(c_result *c_result);

const char *const *c_result_log_ptr(c_result *c_result);

size_t c_result_log_len(c_result *c_result);

#ifdef __cplusplus
} // extern "C"
#endif // __cplusplus
